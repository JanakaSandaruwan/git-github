import ballerina/http;


configurable string str = ?;
configurable string[][] strArr = ?;
configurable string[][][][] janaka = ?;
type Version record {|
    string name;
    string id;
|}

type Resource record {|
     string resourceName;
     string resourceType;
     Version[] version;
|};

type CapabilityStatement record {|
     Resource[] resources;
|};

configurable string greeting = ?;
configurable CapabilityStatement capabilityStatement = ?;


service / on new http:Listener(9090) {

    # A resource for generating greetings
    # + name - the input string name
    # + return - string name with hello message or error
    resource function get greeting(string name) returns string|error {
        // Send a response back to the caller.
        if name is "" {
            return error("name should not be empty ! test comment");
        }
        return str;
    }
}
