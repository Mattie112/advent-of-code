<?php
/** https://adventofcode.com/2019/day/1 */

$total_fuel = 0;
$total_additional_fuel = 0;
if ($file = fopen(__DIR__ . "/day1-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if (!is_numeric($line)) {
            continue;
        }
        $mass = (int)$line;

        // Calculate fuel for this module
        $fuel = floor($mass / 3) - 2;
        $total_fuel += $fuel;

        // Calculate additional fuel for this module (part 2)
        $additional_fuel = $fuel;
        while ($additional_fuel > 0) {
            $additional_fuel = floor($additional_fuel / 3) - 2;
            if ($additional_fuel > 0) {
                $total_additional_fuel += $additional_fuel;
            }
        }

//        echo ($fuel + $total_additional_fuel) . PHP_EOL;
    }
}

echo "Part 1: " . $total_fuel . PHP_EOL;
echo "Part 2: " . ($total_fuel + $total_additional_fuel) . PHP_EOL;


