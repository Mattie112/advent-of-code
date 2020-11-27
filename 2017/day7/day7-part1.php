<?php
/**
 * http://adventofcode.com/2017/day/7
 */

$unsorted_input = [];

if ($file = fopen(__DIR__ . "/day7-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if (empty($line)) {
            continue;
        }

        $unsorted_input[] = $line;
    }
    fclose($file);
}

// First find programs that hold other programs
$programs_that_hold_with_subs = [];
foreach ($unsorted_input as $item) {

    if (strpos($item, "->") !== false) {
        $programs_that_hold_with_subs[] = $item;
    }
}

// Now really split being hold and holding
$programs_being_hold = [];
$programs_that_hold = [];
foreach ($programs_that_hold_with_subs as $item) {
    $matches = [];
    if (preg_match("#-> (.*)#", $item, $matches)) {
        $programs_being_hold[] = $matches[1];
    }
    $matches = [];
    if (preg_match("#([a-zA-Z]*)#", $item, $matches)) {
        $programs_that_hold[] = $matches[1];
    }
}

// Now I can check for the base program by looping through the programs that hold other programs AND programs being hold
foreach ($programs_that_hold as $item) {

    foreach ($programs_being_hold as $prog) {
        if (strpos($prog, $item) !== false) {
            continue 2;
        }
    }

    echo "Found base program '" . $item . "'" . PHP_EOL;
}



