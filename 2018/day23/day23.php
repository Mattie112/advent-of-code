<?php

/**
 * https://adventofcode.com/2018/day/23
 */

$nanobots = [];

if ($file = fopen(__DIR__ . "/day23-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }

        if (preg_match("@pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(-?\d+)@", $line, $matches)) {
            $nanobots[] = ["x" => (int) $matches[1], "y" => (int) $matches[2], "z" => (int) $matches[3], "r" => (int) $matches[4]];
        }
    }
}

$largest_range_bot = findLargestRange($nanobots);
echo "Found largest range: " . $largest_range_bot["r"] . PHP_EOL;

$in_range = findBotsInRange($nanobots, $largest_range_bot);
echo "Part #1: " . count($in_range) . " bots in range of the strongest bot" . PHP_EOL;

function findBotsInRange($nanobots, $strongest)
{
    $in_range = [];
    [$x, $y, $z, $r] = array_values($strongest);

    foreach ($nanobots as $bot) {
        [$bot_x, $bot_y, $bot_z, $bot_r] = array_values($bot);
        if (abs($bot_x - $x) + abs($bot_y - $y) + abs($bot_z - $z) <= $r) {
            $in_range[] = $bot;
        }
    }

    return $in_range;
}

function findLargestRange($nanobots)
{
    $max = 0;
    $max_bot = null;
    foreach ($nanobots as $bot) {
        if ($bot["r"] > $max) {
            $max = $bot["r"];
            $max_bot = $bot;
        }
    }

    return $max_bot;
}
