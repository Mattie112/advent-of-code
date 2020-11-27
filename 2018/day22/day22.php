<?php

/**
 * https://adventofcode.com/2018/day/22
 */

const TARGET = "T";
const MOUTH = "M";
const ROCKY = ".";
const WET = "=";
const NARROW = "|";

$depth = 0;
$target_x = 0;
$target_y = 0;

$cave = [];
$cave_index = [];
$cave[0][0] = MOUTH;
$cave_index[0][0] = 0;
$cave_erosion = [];
if ($file = fopen(__DIR__ . "/day22-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }
        if (preg_match("@depth: (\d+)@", $line, $matches)) {
            $depth = (int) $matches[1];
        }
        if (preg_match("@target: (\d+),(\d+)@", $line, $matches)) {
            $target_x = (int) $matches[1];
            $target_y = (int) $matches[2];
        }
    }
}

echo "Depth: " . $depth . " target: " . $target_x . "," . $target_y . PHP_EOL;

//$cave[$target_y][$target_x] = TARGET;
//$cave_index[$target_y][$target_x] = 0;

for ($y = 0; $y < $depth; $y++) {
    for ($x = 0; $x < $target_x + 1; $x++) {
        $index = 0;
//        if (($x === 0 && $y === 0) || ($x === $target_x && $y === $target_y)) {
//            continue;
//        }
        if ($y === 0) {
            $index = $x * 16807;
        } elseif ($x === 0) {
            $index = $y * 48271;
        } elseif ($x === $target_x && $y === $target_y) {
            $index = 0;
        } else {
            $index = $cave_erosion[$y][$x - 1] * $cave_erosion[$y - 1][$x];
        }
        $cave_index[$y][$x] = $index;

        $erosion = ($index + $depth) % 20183;
        $cave_erosion[$y][$x] = $erosion;
        if ($erosion % 3 === 0) {
            $cave[$y][$x] = ROCKY;
        }
        if ($erosion % 3 === 1) {
            $cave[$y][$x] = WET;
        }
        if ($erosion % 3 === 2) {
            $cave[$y][$x] = NARROW;
        }
    }
}

$cave[0][0] = MOUTH;
$cave[$target_y][$target_x] = TARGET;

printCave($cave, 11);

echo "Part #1: " . calcRisk($cave, $target_x, $target_y) . PHP_EOL;

function calcRisk($cave, $target_x, $target_y)
{
    $risk = 0;
    foreach ($cave as $y => $row) {
        foreach ($row as $x => $cell) {
            if ($x === $target_x && $y === $target_y) {
                return $risk;
            }
            if ($cell === WET) {
                $risk++;
            }
            if ($cell === NARROW) {
                $risk += 2;
            }
        }
    }

    return $risk;
}

function printCave($cave, $maxdepth)
{
    foreach ($cave as $y => $row) {
        if ($y === $maxdepth) {
            return;
        }
        foreach ($row as $x => $cell) {
            echo $cell;
        }
        echo PHP_EOL;
    }
}
