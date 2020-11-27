<?php
/**
 * https://adventofcode.com/2018/day/8
 */

$tree = [];
$input = [];
$part1 = 0;
$part2 = 0;

if ($file = fopen(__DIR__ . "/day8-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }
        $chars = array_map('intval', explode(" ", $line));
        $input = $chars;
        $tree = [buildTree()];
        $part2 = calcPart2(reset($tree));

    }
}

echo "Part 1 " . $part1 . PHP_EOL;
echo "Part 2 " . $part2 . PHP_EOL;

function calcPart2($tree)
{
    $metadata = $tree["metadata"];
    $children = $tree["children"];
    if (count($children) === 0) {
        return array_sum($metadata);
    }

    $calc = 0;
    foreach ($metadata as $index) {
        // Node #1 would be array index 0
        $index--;
        if (isset($children[$index])) {
            $calc += calcPart2($children[$index]);
        }
    }

    return $calc;
}

function buildTree()
{
    global $input;
    global $part1;
    $header_child_amount = array_shift($input);
    $header_meta_amount = array_shift($input);

    $children = [];

    for ($i = 0; $i < $header_child_amount; $i++) {
        $children[] = buildTree();
    }

    $metadata = [];
    for ($j = 0; $j < $header_meta_amount; $j++) {
        $value = array_shift($input);
        $metadata[] = $value;
        // No need to loop over all elements when finished we will simply calcualte the checksum here already
        $part1 += $value;
    }

    return ["children" => $children, "metadata" => $metadata];
}
