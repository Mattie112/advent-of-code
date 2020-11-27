<?php

/**
 * https://adventofcode.com/2018/day/14
 */
error_reporting(E_ALL);
$input = "165061";
// Tis is what we start with
$recepies = "37";
// Test data
$elves_goal = 9;
$elves_goal = 5;
$elves_goal = 18;
$elves_goal = 2018;

$search_string = "51589";
$search_string = "01245";
$search_string = "92510";
$search_string = "59414";

// Puzzle input
$elves_goal = $input;
$search_string = $input;
$search_len = strlen($search_string);

$goal = (int) $elves_goal + 10;

$elve_positions = [0, 1];
$part1 = false;
$part2 = false;
//debugElves($recepies, $elve_positions);

while (!$part1 || !$part2) {
    // Calculate new recepies
    $try = (int) $recepies[$elve_positions[0]];
    $try += (int) $recepies[$elve_positions[1]];
    $recepies .= $try;

    $new_amount = strlen($recepies);

    foreach ($elve_positions as &$position) {
        $position = ($position + $recepies[$position] + 1) % $new_amount;
    }
    unset($position);

    if (!$part1 && $new_amount === $goal) {
        $part1 = substr($recepies, -10);
    }

    // Try to save some time by only looking at a part of the string with 2 extra characters to be sure
    if (!$part2 && strpos(substr($recepies, -$search_len - 2), $search_string) !== false) {
        $end = strpos($recepies, $search_string);
        $part2 = strlen(substr($recepies, 0, strlen($recepies) - strlen($search_string) - 1));
    }

//    debugElves($recepies, $elve_positions);
}

echo "Part #1: " . $part1 . PHP_EOL;
echo "Part #2: " . $part2 . PHP_EOL;

function debugElves($recepies, $elve_positions)
{
    $new_amount = strlen($recepies);
    foreach ($recepies as $i => $iValue) {
        if ($i === $elve_positions[0]) {
            echo "(" . $iValue . ")";
        } elseif ($i === $elve_positions[1]) {
            echo "[" . $iValue . "]";
        } else {
            echo " " . $iValue . " ";
        }
    }
    echo PHP_EOL;
}
