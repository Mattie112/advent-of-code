<?php
/** https://adventofcode.com/2019/day/5 */
const MODE_POS = 0;
const MODE_IMMIDIATE = 1;
//$input = file_get_contents(__DIR__ . "/day2-input.txt");
//$input = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,10,23,1,23,6,27,1,6,27,31,1,13,31,35,1,13,35,39,1,39,13,43,2,43,9,47,2,6,47,51,1,51,9,55,1,55,9,59,1,59,6,63,1,9,63,67,2,67,10,71,2,71,13,75,1,10,75,79,2,10,79,83,1,83,6,87,2,87,10,91,1,91,6,95,1,95,13,99,1,99,13,103,2,103,9,107,2,107,10,111,1,5,111,115,2,115,9,119,1,5,119,123,1,123,9,127,1,127,2,131,1,5,131,0,99,2,0,14,0";
//$input = "3,225,1,225,6,6,1100,1,238,225,104,0,1002,114,46,224,1001,224,-736,224,4,224,1002,223,8,223,1001,224,3,224,1,223,224,223,1,166,195,224,1001,224,-137,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,1001,169,83,224,1001,224,-90,224,4,224,102,8,223,223,1001,224,2,224,1,224,223,223,101,44,117,224,101,-131,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,1101,80,17,225,1101,56,51,225,1101,78,89,225,1102,48,16,225,1101,87,78,225,1102,34,33,224,101,-1122,224,224,4,224,1002,223,8,223,101,7,224,224,1,223,224,223,1101,66,53,224,101,-119,224,224,4,224,102,8,223,223,1001,224,5,224,1,223,224,223,1102,51,49,225,1101,7,15,225,2,110,106,224,1001,224,-4539,224,4,224,102,8,223,223,101,3,224,224,1,223,224,223,1102,88,78,225,102,78,101,224,101,-6240,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1107,226,677,224,102,2,223,223,1006,224,329,101,1,223,223,1108,226,677,224,1002,223,2,223,1005,224,344,101,1,223,223,8,226,677,224,102,2,223,223,1006,224,359,1001,223,1,223,1007,226,677,224,1002,223,2,223,1005,224,374,101,1,223,223,1008,677,677,224,1002,223,2,223,1005,224,389,1001,223,1,223,1108,677,226,224,1002,223,2,223,1006,224,404,1001,223,1,223,1007,226,226,224,1002,223,2,223,1005,224,419,1001,223,1,223,1107,677,226,224,1002,223,2,223,1006,224,434,101,1,223,223,108,677,677,224,1002,223,2,223,1005,224,449,1001,223,1,223,1107,677,677,224,102,2,223,223,1005,224,464,1001,223,1,223,108,226,226,224,1002,223,2,223,1006,224,479,1001,223,1,223,1008,226,226,224,102,2,223,223,1005,224,494,101,1,223,223,108,677,226,224,102,2,223,223,1005,224,509,1001,223,1,223,8,677,226,224,1002,223,2,223,1006,224,524,101,1,223,223,7,226,677,224,1002,223,2,223,1006,224,539,101,1,223,223,7,677,226,224,102,2,223,223,1006,224,554,1001,223,1,223,7,226,226,224,1002,223,2,223,1006,224,569,101,1,223,223,107,677,677,224,102,2,223,223,1006,224,584,101,1,223,223,1108,677,677,224,102,2,223,223,1006,224,599,1001,223,1,223,1008,677,226,224,1002,223,2,223,1005,224,614,1001,223,1,223,8,677,677,224,1002,223,2,223,1006,224,629,1001,223,1,223,107,226,677,224,1002,223,2,223,1006,224,644,101,1,223,223,1007,677,677,224,102,2,223,223,1006,224,659,101,1,223,223,107,226,226,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226";
$input = "1002,4,3,4,33";
$program = explode(",", $input);
$part2_expected = 19690720;
$program[1] = 12;
$program[2] = 02;
echo "Part 1: " . calculate($program) . PHP_EOL;
//// Now find the noun, simply iterate till we have a value that is to big
//$value = calculate($program);
//while ($value < $part2_expected) {
//    $program[1]++;
//    $value = calculate($program);
//}
//// Now subtract one: this is our noun
//$program[1]--;
//// Now find the verb by doing the same, but now we look for an exact match
//$value = calculate($program);
//while ($value !== $part2_expected) {
//    $program[2]++;
//    $value = calculate($program);
//}
//echo "Part 2: " . (100 * $program[1] + $program[2]) . " (noun: " . $program[1] . " - verb: " . $program[2] . ")" . PHP_EOL;
function calculate(array $program)
{
    $index = 0;
    while (true) {
        // check parameter 1
        switch (floor($program[$index] / 100) % 10) {
            case MODE_POS:
                $param1 = $program[$program[$index + 1]];
                break;
            case MODE_IMMIDIATE:
                $param1 = $program[$index + 1];
                break;
            default:
                die("Unkown parameter 1" . PHP_EOL);
        }
        if ($param1 === null) {
            $param1 = MODE_POS;
        }

        // check parameter 2
        switch (floor($program[$index] / 1000) % 10) {
            case MODE_POS:
                $param2 = $program[$program[$index + 1]];
                break;
            case MODE_IMMIDIATE:
                $param2 = $program[$index + 1];
                break;
            default:
                die("Unkown parameter 2" . PHP_EOL);
        }
        if ($param2 === null) {
            $param1 = MODE_POS;
        }

        // check parameter 3
        switch (floor($program[$index] / 10000) % 10) {
            case MODE_POS:
                $param3 = $program[$program[$index + 1]];
                break;
            case MODE_IMMIDIATE:
                $param3 = $program[$index + 1];
                break;
            default:
                die("Unkown parameter 3" . PHP_EOL);
        }
        if ($param3 === null) {
            $param3 = MODE_POS;
        }

        $opcode = $program[$index] % 100;
        switch ($opcode) {
            case 1:
                // Add numbers (1, 2) and store into (3)
                $number1 = $program[$program[$index + 1]];
                $number2 = $program[$program[$index + 2]];
                $program[$program[$index + 3]] = $number1 + $number2;
                $index += 4;
                break;
            case 2:
                // Multiply numbers (1, 2) and store into (3)
                $number1 = $program[$program[$index + 1]];
                $number2 = $program[$program[$index + 2]];
                $program[$program[$index + 3]] = $number1 * $number2;
                $index += 4;
                break;
            case 3:
                // Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.

                break;
            case 4:
                // Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.

                break;
            case 9:
                // hack to prefent 2 digits as opcode
            case 99:
                return $program[0];
            default:
                // e.g. unknown upcode
                die("Unknown opcode " . $opcode . PHP_EOL);
        }
    }
    return null;
}
