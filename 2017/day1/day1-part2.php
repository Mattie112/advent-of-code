<?php
/**
 * http://adventofcode.com/2017/day/1
 */

$input = trim(file_get_contents(__DIR__ . "/day1-input.txt"));
$input = str_split($input);
$total_count = count($input);
$forward = $total_count / 2;
$answer = 0;

foreach ($input as $count => $number) {

    $forward = $count + ($total_count / 2);
    if ($forward > (count($input) - 1)) {
        $forward -= count($input);
    }

    if ($number === $input[$forward]) {
        $answer += $number;
    }
}

echo $answer . PHP_EOL;
