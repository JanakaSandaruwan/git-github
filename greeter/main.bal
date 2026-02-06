import ballerina/io;
import ballerina/time;
import  ballerina/http as _;

// type Coord record {
//     int x;
//     int y;
// };

type flexType string|int;

configurable string greeting = "Hello, Test World!";
// configurable int[] | string[] repeatCount = ?;
// configurable boolean isEnabled = true;
configurable float version = 1.0;
// configurable map<string> settings = ?;
// configurable Coord position = { x: 10, y: 20 };
// configurable flexType dynamicValue = "dynamic";
// configurable int[] numbers = [1, 2, 3, 4, 5];
// configurable map<int[]> multiDimensionalMap = ?;
// configurable decimal decimalValue = 3.14;
// configurable string? optionalGreeting = "Optional Hello";
// configurable [string, int, boolean] person = ["Alice", 30, true];
// configurable json jsonData = {
//     "name": "Bob",
//     "age": 25,
//     "isActive": true
// };

configurable float pi = 3;
configurable string num = ?;
configurable time:TimeOfDay lunchFeedbackStartTime = {hour: 11, minute: 0, second: 0};

public function main() {
    io:println("Configurable value of pi: ", lunchFeedbackStartTime);
}
