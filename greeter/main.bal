import ballerina/http;

# A Ballerina service representing a network-accessible API
# bound to port `9090`.

// type Resource record {|
//     string resourceName;
//     string resourceType;
// |};

// type CapabilityStatement record {|
//     Resource[] resources;
// |};

configurable string greeting = ?;
// configurable CapabilityStatement capabilityStatement2 = ?;
// configurable map<string[]> kamal = ?;
configurable string arr = ?;
configurable string str = ?;
// configurable string[] str1 = ?;
// configurable string str = ?;
// configurable string str3 = ?;
// configurable Resource resourceRecord = ?;


service / on new http:Listener(9090) {

    # A resource for generating greetings
    # + name - the input string name
    # + return - string name with hello message or error
    resource function get greeting(string name) returns string|error {
        // Send a response back to the caller.
        if name is "" {
            return error("name should not be empty ! test comment");
        }
        return name;
    }
}
