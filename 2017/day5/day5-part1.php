<?php
/**
 * http://adventofcode.com/2017/day/5
 */

$answer = 0;

$steps = [];

// Read the file and store in array
if ($file = fopen(__DIR__ . "/day5-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if ($line === "") {
            continue;
        }

        $steps[] = (int) $line;
    }
    fclose($file);
}

$position = 0;
$total_steps = count($steps) - 1;
$go = true;

while ($go) {
    if ($position > $total_steps) {
        $go = false;
        continue;
    }

    $step = &$steps[$position];
    $answer++;

    if ($step === 0) {
        $step++;
        continue;
    }

    // As the item in the array is already negative we add the steps in all cases
    $position += $step;
    $step++;
}

echo $answer . PHP_EOL;
