<?php
/**
 * http://adventofcode.com/2017/day/2
 */

$answer = null;

if ($file = fopen(__DIR__ . "/day2-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        // Accept tabs and space
        $line = preg_split('/\s+/', $line);

        foreach ($line as $key => $number) {
            foreach ($line as $devision_key => $devision_number) {
                // Don't devide by itself
                if ($key === $devision_key) {
                    continue;
                }

                if (is_int($number / $devision_number)) {
                    $answer += $number / $devision_number;
                }
            }
        }
    }
    fclose($file);
}

echo $answer . PHP_EOL;
