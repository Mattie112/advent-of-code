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

        $max = null;
        $min = null;

        foreach ($line as $number) {
            $number = (int) $number;
            if (!$max || $number > $max) {
                $max = $number;
            }

            if (!$min || $number < $min) {
                $min = $number;
            }
        }

        $diff = $max - $min;
        $answer += $diff;

    }
    fclose($file);
}

echo $answer . PHP_EOL;
