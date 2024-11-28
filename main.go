package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Get the server URL from environment variable
	url := os.Getenv("SERVER_URL")
	if url == "" {
		log.Fatal("Environment variable SERVER_URL is not set")
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

	// Create a custom HTTP transport using the TLS config
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Create an HTTP client with the transport
	client := &http.Client{
		Transport: transport,
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Set the custom Host header if provided
	if customHostHeader != "" {
		req.Header.Set("Host", customHostHeader)
		log.Printf("Using custom Host header: %s", customHostHeader)
	}

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make HTTPS request: %v", err)
	}
	defer resp.Body.Close()

	// Print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response from server:\n%s\n", body)
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
