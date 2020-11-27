<?php
/**
 * http://adventofcode.com/2017/day/8
 */

$registers = [];
$highest = 0;

// Read the file and store in array
if ($file = fopen(__DIR__ . "/day8-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if ($line === "") {
            continue;
        }

        list($register, $operator, $value, , $if_register, $if_operator, $if_value) = preg_split('/\s+/', $line);

        // Make sure both registers exists before doing anything else
        if (!isset($registers[$register])) {
            $registers[$register] = 0;
        }

        if (!isset($registers[$if_register])) {
            $registers[$if_register] = 0;
        }

        // Eval FTW!
        $eval_str = "return " . $registers[$if_register] . $if_operator . $if_value . ";";
        if (isset($registers[$if_register]) && eval($eval_str)) {
            if ($operator === "inc") {
                $registers[$register] += $value;
            } else {
                $registers[$register] -= $value;
            }
        }

        $highest = max($highest, $registers[$register]);

    }
    fclose($file);
}

echo "Found highest register with value: " . max($registers) . PHP_EOL;
echo "Memory needed to allocate: " . $highest . PHP_EOL;
