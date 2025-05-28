// import ballerina/http;

import ballerina/io;
import ballerina/time;

configurable string str = ?;
configurable string abc = ?;
configurable string ba = ?;

public function main() {
    io:println("Hello, World!");
    time:sleep(3600 * 1000); // sleep takes milliseconds
    io:println("Woke up after 1 hour.");
}
