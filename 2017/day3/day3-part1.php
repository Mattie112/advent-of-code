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
    $grid[$y][$x] = $last_value;
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

// We want to know the steps for field we'll need to find
krsort($grid);
foreach ($grid as $x2 => $y_arr) {
    ksort($y_arr);
    foreach ($y_arr as $y2 => $yvalue) {
        if ($yvalue === $field_to_find) {
            echo "At x: " . $x2 . " and y: " . $y2 . "  we have found " . $yvalue . PHP_EOL;
            $answer = abs($x2) + abs($y2);
        }
    }
}

echo "Answer: " . $answer . PHP_EOL;

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
