<?php
/**
 * http://adventofcode.com/2017/day/7
 */

$tree = [];
$unsorted_input = [];

// Well I was in no mood to change the part-1 code to the new array-structure so keeping that intact
$unparsed_input = [];

if ($file = fopen(__DIR__ . "/day7-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if (empty($line)) {
            continue;
        }

        $unparsed_input[] = $line;

        $program = "";
        if (preg_match("#([a-zA-Z]*)#", $line, $matches)) {
            $program = $matches[1];
        }

        $children = [];
        if (strpos($line, "->") !== false) {
            if (preg_match("#-> (.*)#", $line, $matches)) {
                $children = $matches[1];
                $children = explode(", ", $children);
            }
        }

        $weight = "?";
        if (preg_match("#([\d]+)#", $line, $matches)) {
            $weight = $matches[1];
        }

        $unsorted_input[$program] = [
            "weight" => $weight,
            "children" => $children,
        ];

    }
    fclose($file);
}

// First find programs that hold other programs
$programs_that_hold_with_subs = [];
foreach ($unparsed_input as $item) {

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
$base_program = "";
foreach ($programs_that_hold as $item) {
    foreach ($programs_being_hold as $prog) {
        if (strpos($prog, $item) !== false) {
            continue 2;
        }
    }

    $base_program = $item;
    echo "Found base program '" . $item . "'" . PHP_EOL;
}

// Now that we know the base program we can build the tree (yeah finally)

$tree[$base_program] = [];

$something = findChildren($base_program, $unsorted_input);
$tree[$something[0]] = $something[1];

function findChildren($base, &$heep)
{
    // Find the entire program
    $program = $heep[$base];

    if (count($program["children"]) === 0) {
        unset($heep[$base]);
        $program["total_weight"] = $program["weight"];

        return [$base, $program];
    }

    $children = [];
    $children_weight = 0;
    foreach ($program["children"] as $child) {
        list($name, $child_program) = findChildren($child, $heep);
        $children[$name] = $child_program;
        $children_weight += $child_program["total_weight"];
    }

    $program["children"] = $children;
    $program["total_weight"] = $children_weight + $program["weight"];

    unset($heep[$base]);

    return [$base, $program];
}

// Now the tree is completed and we can work on the weight
findWrongWeight($tree[$base_program], $base_program, 0);

function findWrongWeight($tree, $program_name, $correction)
{
    $weights = [];
    foreach ($tree["children"] as $key => $children) {
        $weights[$children["total_weight"]][$key] = $children;
    }

    // If we have the same weight we finally know it!
    if (count($weights) === 1) {
        $incorrect_weight = $tree["weight"];
        echo "Program '" . $program_name . "' needs to be: " . ($incorrect_weight + $correction) . PHP_EOL;
        die();
    }

    $correction = array_search(max($weights), $weights) - array_search(min($weights), $weights);
    $wrong = $weights[array_search(min($weights), $weights)];

    findWrongWeight(reset($wrong), key($wrong), $correction);
}
