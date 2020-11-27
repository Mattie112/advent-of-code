<?php
/**
 * http://adventofcode.com/2017/day/1
 */

$input = trim(file_get_contents(__DIR__ . "/day1-input.txt"));
$input = str_split($input);
// The last item must be matched with the first so adding it here
$input[] = $input[0];
$total_count = count($input);
$prev = null;
$answer = 0;

foreach ($input as $count => $number) {
    if ($prev && $prev === $number) {
        $answer += (int) $number;
    }

    $prev = $number;
}

echo $answer . PHP_EOL;
