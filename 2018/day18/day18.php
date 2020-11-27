<?php
/**
 * https://adventofcode.com/2018/day/18
 */
const TREE = "|";
const LUMBERYARD = "#";
const OPEN_GROUND = ".";

$land = [];
if ($file = fopen(__DIR__ . "/day18-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }
        $land[] = str_split($line);
    }
}

$part1 = false;
$part2 = false;
$prev_amount = 0;
$previous_states = [];
for ($i = 0; $i < 1000000000; $i++) {
    $new_land = $land;
    foreach ($land as $y => $line) {
        foreach ($line as $x => $tile) {
            // Find adjacent tiles
            $find_tiles = [
                $land[$y][$x - 1] ?? " ",
                $land[$y][$x + 1] ?? " ",
                $land[$y - 1][$x] ?? " ",
                $land[$y + 1][$x] ?? " ",
                $land[$y + 1][$x + 1] ?? " ",
                $land[$y + 1][$x - 1] ?? " ",
                $land[$y - 1][$x + 1] ?? " ",
                $land[$y - 1][$x - 1] ?? " ",
            ];
            $find_tiles_values = array_count_values($find_tiles);
            switch ($tile) {
                case TREE:
                    if ($find_tiles_values[LUMBERYARD] >= 3) {
                        $new_land[$y][$x] = LUMBERYARD;
                    }
                    break;
                case LUMBERYARD:
                    if ($find_tiles_values[LUMBERYARD] >= 1 && $find_tiles_values[TREE] >= 1) {
                        $new_land[$y][$x] = LUMBERYARD;
                    } else {
                        $new_land[$y][$x] = OPEN_GROUND;
                    }
                    break;
                case OPEN_GROUND:
                    if ($find_tiles_values[TREE] >= 3) {
                        $new_land[$y][$x] = TREE;
                    }
                    break;
            }
        }
    }
    if ($i === 10) {
        $part1 = $land;
    }
    $land = $new_land;

    $tmp = array_count_values(array_merge(...$land));
    $amount = $tmp[TREE] * $tmp[LUMBERYARD];

    $prev_step = array_search($land, $previous_states, true);
    // If we have found a previous answer add the amount of steps to $i but never more then 1000000000
    if ($prev_step !== false) {
        $i = (floor((1000000000 - $i) / ($i - $prev_step)) * ($i - $prev_step) + $i);
    }

    $buffer = "After " . $i . " minutes:" . PHP_EOL;
    $buffer = printLand($land, $buffer);
    $buffer .= PHP_EOL;
    echo $buffer;
    $prev_amount = $amount;
    $previous_states[] = $land;
}

$arr_val = array_count_values(array_merge(...$part1));
$arr_val2 = array_count_values(array_merge(...$land));
echo "Part #1: " . $arr_val[TREE] . " wooded acres and " . $arr_val[LUMBERYARD] . " lumberyards. Answer: " . $arr_val[TREE] * $arr_val[LUMBERYARD] . PHP_EOL;
echo "Part #2: " . $arr_val2[TREE] . " wooded acres and " . $arr_val2[LUMBERYARD] . " lumberyards. Answer: " . $arr_val2[TREE] * $arr_val2[LUMBERYARD] . PHP_EOL;

function printLand($land, $buffer)
{
    foreach ($land as $l) {
        foreach ($l as $ll) {
            $buffer .= $ll;
        }
        $buffer .= PHP_EOL;
    }

    return $buffer;
}

