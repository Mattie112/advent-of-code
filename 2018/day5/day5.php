<?php
/**
 * https://adventofcode.com/2018/day/5
 */

$input = trim(file_get_contents(__DIR__ . "/day5-input.txt"));

// Prepare the ranges we want to remove somehow working with strreplace is way way way quicker then looping through all letters
// (or my skills are not good enough that is also a possibility).
$letter_range = range('a', 'z');
$letter_range_upper = range('A', 'Z');

// For part 2
$improve_length = PHP_INT_MAX;
$improve_letter = "";

// First we do part 2 as my part 1 code modifies the $input
foreach ($letter_range as $index_part2 => $letter_part2) {
    // Foreach letter copy the input again, remove ALL instances from that letter (caseINsensitive) and execute the part 1 calculation
    // Whenever we have 'less' length store the amount and the related letter
    $tmp_input = $input;
    $tmp_input = str_ireplace($letter_part2, "", $tmp_input);

    $tmp_input = calc($tmp_input, $letter_range, $letter_range_upper);

    if (strlen($tmp_input) < $improve_length) {
        $improve_length = strlen($tmp_input);
        $improve_letter = $letter_part2;
    }
}

// This is part 1
$input = calc($input, $letter_range, $letter_range_upper);

echo "Answer part 1: " . strlen($input) . PHP_EOL;
echo "When you remove '" . $improve_letter . "' you get the best result (" . $improve_length . " length)";

// Shared for both days, this will give the answer for part 1, simply replace aA and Aa by nothing (and then for a-z A-Z)
function calc($input, $letter_range, $letter_range_upper)
{
    while (true) {
        $replaced_count = 0;
        foreach ($letter_range as $index => $letter) {
            $input = str_replace($letter . $letter_range_upper[$index], "", $input, $temp);
            $input = str_replace($letter_range_upper[$index] . $letter, "", $input, $temp2);
            $replaced_count = max($replaced_count, $temp, $temp2);
        }
        // If we failed to replace anything we are done
        if ($replaced_count === 0) {
            return $input;
        }
    }
    return null;
}
