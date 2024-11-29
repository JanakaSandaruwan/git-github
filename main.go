package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

func main() {
	// Get the server URL from environment variable
	targetUrl := os.Getenv("TARGET_SERVER_URL")
	targetURL, err := url.Parse(targetUrl)
	if err != nil {
		log.Fatalf("Invalid TARGET_SERVER_URL: %v", err)
	}

	// Get the CA certificate path from environment variable
	caCertPath := os.Getenv("CA_CERT_PATH")
	if caCertPath == "" {
		log.Fatal("Environment variable CA_CERT_PATH is not set")
	}

	// Get options to disable TLS and host verification
	disableTLSVerification := getEnvAsBool("DISABLE_TLS_VERIFICATION", false)
	disableHostVerification := getEnvAsBool("DISABLE_HOST_VERIFICATION", false)

	// Get the custom Host header value
	customHostHeader := os.Getenv("CUSTOM_HOST_HEADER")

	// Load the server's CA certificate
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}

	// Create a certificate pool and append the CA cert
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Fatal("Failed to append CA certificate to pool")
	}

	// Configure TLS settings
	tlsConfig := &tls.Config{
		// RootCAs:    caCertPool,
		ServerName: customHostHeader,
	}
	if disableTLSVerification {
		log.Println("TLS verification is disabled")
		tlsConfig.InsecureSkipVerify = true
	}
	if disableHostVerification {
		log.Println("Host verification is disabled")
		tlsConfig.VerifyPeerCertificate = nil
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Proxying request: %s %s", r.Method, r.URL.String())

		// Rewrite the request's Host header to the target server
		r.Host = customHostHeader

		// Forward the request to the target server
		proxy.ServeHTTP(w, r)
	})

	// Start the proxy server
	port := os.Getenv("PROXY_PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}
	log.Printf("Starting proxy server on port %s, forwarding to %s", port, targetURL)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// getEnvAsBool fetches an environment variable as a boolean with a default fallback
func getEnvAsBool(envVar string, defaultVal bool) bool {
	valStr := os.Getenv(envVar)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.ParseBool(valStr)
	if err != nil {
		log.Printf("Invalid boolean value for %s: %v. Using default: %v", envVar, valStr, defaultVal)
		return defaultVal
	}
	return val
}
