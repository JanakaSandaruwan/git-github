// import ballerina/http;


configurable string str = ?;
configurable boolean flag = ?;
configurable string[][] strArr = ?;
configurable string janaka = "janka"; // test ned
type Version record {|
    string name = "janaka";
    string id;
|};

type Resource record {|
     string resourceName;
     string resourceType;
     Version[] version;
|};

type CapabilityStatement record {|
     Resource[] resources;
     string name;
|};

configurable string greeting = ?;
configurable CapabilityStatement capabilityStatement = ?;
configurable Version parentVersion = { name: "defaultName", id: "defaultId" }; // test
