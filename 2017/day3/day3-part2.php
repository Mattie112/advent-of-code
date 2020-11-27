<?php
/**
 * http://adventofcode.com/2017/day/3
 */

// My puzzle input
$field_to_find = 277678;

$answer = null;
$grid = [];
$last_value = 1;
$x = 0;
$y = 0;
$right = 0;
$up = 0;
$left = 0;
$down = 0;
$step_right_up = 1;
$step_left_down = 2;

while ($last_value <= $field_to_find) {

    // New empty space to be filled, determine the value

    $permutations = [[0, 1], [1, 1], [1, 0], [1, -1], [0, -1], [-1, -1], [-1, 0], [-1, 1]];
    $value_to_fill = 0;
    foreach ($permutations as $perm) {
        if (isset($grid[$y + $perm[0]][$x + $perm[1]])) {
            $value_to_fill += $grid[$y + $perm[0]][$x + $perm[1]];
        }
    }
    if ($x === 0 && $y === 0) {
        $value_to_fill = 1;
    }
    $grid[$y][$x] = $value_to_fill;
    if ($value_to_fill > $field_to_find) {
        echo "Found the answer: " . $value_to_fill . PHP_EOL;
        draw($grid);
        die();
    }
    $last_value++;

    if ($right < $step_right_up) {
        $right++;
        $x++;
        draw($grid);
        continue;
    }

    if ($up < $step_right_up) {
        $up++;
        $y++;
        draw($grid);
        continue;
    }

    if ($left < $step_left_down) {
        $left++;
        $x--;
        draw($grid);
        continue;
    }

    if ($down < $step_left_down) {
        $down++;
        $y--;
        draw($grid);
        continue;
    }

    // Oh we are done with this 'circle' lets start over
    $last_value--;
    $right = 0;
    $up = 0;
    $left = 0;
    $down = 0;
    $step_left_down++;
    $step_left_down++;
    $step_right_up++;
    $step_right_up++;

    echo "*******************LOOP DONE*********************" . PHP_EOL;
    echo "New step right up: " . $step_right_up . PHP_EOL;
    echo "New step left down: " . $step_left_down . PHP_EOL;
    echo "Cursor pos: x:" . $x . " y:" . $y . PHP_EOL;
    echo "*******************LOOP DONE*********************" . PHP_EOL;
}

function draw($grid)
{
    // Enable to spam your console
    return;
    echo str_repeat("-", 50) . PHP_EOL;
    // See if we can print the grid
    krsort($grid);
    foreach ($grid as $x2 => $y_arr) {
        ksort($y_arr);
        foreach ($y_arr as $y2) {
            echo "$y2 \t";
        }
        echo PHP_EOL;
    }
}
