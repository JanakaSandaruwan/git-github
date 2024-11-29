import ballerina/http;
import ballerina/io;


configurable string Endpoint = "https://localhost";
configurable string HostName = "service1.example.com";
configurable string certPath = "/tmp/service1.crt";
configurable boolean verifyHost = false;
configurable boolean enableSecurity = false;

public function main() returns error? {
   http:Client backendEP = check new(Endpoint, config = {
        secureSocket: {
            enable: enableSecurity,
            verifyHostName: verifyHost,
            cert: certPath
        }
    });

    http:Response backendResponse = check backendEP->get("/", headers = {
        "Host" : HostName
    });
    io:println(backendResponse.getTextPayload());
}
