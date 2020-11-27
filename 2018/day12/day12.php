<?php
/**
 * https://adventofcode.com/2018/day/12
 */

$initial = "";
$notes = [];

if ($file = fopen(__DIR__ . "/day12-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }

        if (preg_match("@initial state: (.*)@", $line, $matches)) {
            $initial = $matches[1];
            $initial = str_split($initial);
        }

        if (preg_match("@(.*) => (.)@", $line, $matches)) {
            $notes[$matches[1]] = $matches[2];
        }
    }
}

// Start by adding the 'left' pots
$pots = [];
$pots[-3] = ".";
$pots[-2] = ".";
$pots[-1] = ".";

// Setup the initial pots
foreach ($initial as $item) {
    if ($item === "#") {
        $pots[] = "#";
    } else {
        $pots[] = ".";
    }
}

// Now add the 'right' pots
$pots[] = ".";
$pots[] = ".";
$pots[] = ".";

$next_iter_pots = $pots;
$generations = 50000000000 - 1;
$count = 0;
$number_count = 0;
$part1 = 0;
$part2 = 0;
$prev_answer = 0;
$prev_plants = "";

for ($i = 0; $i <= $generations; $i++) {
    foreach ($pots as $index => $pot) {
        // We simply use the string to check the nodes
        $str_to_check = "";
        $str_to_check .= $pots[$index - 2] ?? ".";
        $str_to_check .= $pots[$index - 1] ?? ".";
        $str_to_check .= $pots[$index];
        $str_to_check .= $pots[$index + 1] ?? ".";
        $str_to_check .= $pots[$index + 2] ?? ".";

        $next_iter_pots[$index] = $notes[$str_to_check] ?? ".";
    }

    // Add more pots to the left when needed
    reset($next_iter_pots);
    $first_key = key($next_iter_pots);
    if ($next_iter_pots[$first_key + 2] === "#") {
        $next_iter_pots[$first_key - 1] = ".";
        ksort($next_iter_pots);
    }

    // As we have negative indexes we'll need to know the last element
    end($next_iter_pots);
    $last_key = key($next_iter_pots);
    // If we end with a plant add more pots (that is how I read the example)
    if ($next_iter_pots[$last_key - 2] === "#") {
        $next_iter_pots[] = ".";
    }

    // Store the next iteration into the current iteration
    $pots = $next_iter_pots;
    $answer = calcAnswer($pots);
    echo implode("", $pots) . PHP_EOL;

    if ($i === 20) {
        $part1 = $answer;
    }

    // In the end it seems that we are only "moving" our plants and not adding new
    // so if we take only the pots that contain a plant and check that we know if thare are any differences
    $first_plant = array_search("#", $pots, true);
    $plants = substr(implode("", $pots), $first_plant);
    if ($prev_plants === $plants && $i > 20) {
        // If not then we are not doing anything usefull and we can calculate the answer without looping YEAH!
        // So: our current answer + the diff between the last to (that is added each generation) but substract the current generation as that is already in our $answer
        $part2 = $answer + (($answer - $prev_answer) * ($generations - $i));
        break;
    }

    $prev_plants = $plants;
    $prev_answer = $answer;
}

function calcAnswer($pots)
{
    $number_count = 0;
    foreach ($pots as $index => $pot) {
        if ($pot === "#") {
            $number_count += $index;
        }
    }

    return $number_count;
}


echo PHP_EOL . "Part #1: " . $part1 . PHP_EOL;
echo PHP_EOL . "Part #2: " . $part2 . PHP_EOL;
