<?php
/** https://adventofcode.com/2019/day/2 */

$input = file_get_contents(__DIR__ . "/day2-input.txt");
$program = explode(",", $input);
$part2_expected = 19690720;

$program[1] = 12;
$program[2] = 02;

echo "Part 1: " . calculate($program) . PHP_EOL;

// Now find the noun, simply iterate till we have a value that is to big
$value = calculate($program);
while ($value < $part2_expected) {
    $program[1]++;
    $value = calculate($program);
}
// Now subtract one: this is our noun
$program[1]--;

// Now find the verb by doing the same, but now we look for an exact match
$value = calculate($program);
while ($value !== $part2_expected) {
    $program[2]++;
    $value = calculate($program);
}

echo "Part 2: " . (100 * $program[1] + $program[2]) . " (noun: " . $program[1] . " - verb: " . $program[2] . ")" . PHP_EOL;

function calculate(array $program)
{
    $index = 0;
    while (true) {
        $opcode = (int)$program[$index];
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
            case 99:
                return $program[0];
            default:
                // e.g. unknown upcode
                die("Unknown opcode " . $opcode);
        }
    }
    return null;
}
