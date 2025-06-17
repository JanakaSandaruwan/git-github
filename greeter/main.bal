import ballerina/io;
import ballerina/http as _;

type Coord record {
    int x;
    int y;
};

type flexType string|int;

configurable string greeting = "Hello, World!";
configurable map<string> headers = ?;
configurable int repeatCount = 1;
configurable boolean isEnabled = true;
configurable float version = 1.0;
configurable map<string> settings = {
    "theme": "dark",
    "language": "en"
};
configurable Coord position = { x: 10, y: 20 };
configurable flexType dynamicValue = "dynamic";
configurable int[] numbers = [1, 2, 3, 4, 5];
configurable map<int[]> multiDimensionalMap = {
    "1" : [10, 20],
    "2" : [30, 40]
};
configurable decimal decimalValue = 3.14;
configurable string? optionalGreeting = "Optional Hello";
configurable [string, int, boolean] person = ["Alice", 30, true];
configurable json jsonData = {
    "name": "Bob",
    "age": 25,
    "isActive": true
};

public function main() {
    io:println("Hello, World!");
}
