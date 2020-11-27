<?php
/** https://adventofcode.com/2019/day/3 */

const UP = "U";
const LEFT = "L";
const RIGHT = "R";
const DOWN = "D";

$input = file_get_contents(__DIR__."/day3-input.txt");
//$input = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
//U98,R91,D20,R16,D67,R40,U7,R15,U6,R7";
$wires = explode("\n", $input);

$grid [0][0] = "O";
$debug_grid [0][0] = "O";
$wire_count = [];
$intersection_steps = PHP_INT_MAX;
$intersections = [];
$intersections_by_pos = [];

foreach ($wires as $wire_id => $wire_actions) {
    $xpos = 0;
    $ypos = 0;
    $wire_count[$wire_id] = 0;
    $wire_actions = explode(',', $wire_actions);
    foreach ($wire_actions as $wire_action) {
        $wire_action = str_split($wire_action);
        $direction = $wire_action[0];
        unset($wire_action[0]);
        $distance = (int)implode($wire_action);

        switch ($direction) {
            case UP:
                for ($i = 0; $i < $distance; $i++) {
                    $xpos++;
                    $wire_count[$wire_id]++;
                    if (hasIntersection($debug_grid, $xpos, $ypos, $wire_id)) {
                        $intersection_steps = min($intersection_steps, $wire_count[$wire_id]);
                    }
                    [$debug_grid, $intersections] = drawCharacter($debug_grid, $xpos, $ypos, $direction, $wire_id, $intersections);
                }
                break;
            case LEFT:
                for ($i = 0; $i < $distance; $i++) {
                    $ypos--;
                    $wire_count[$wire_id]++;
                    if (hasIntersection($debug_grid, $xpos, $ypos, $wire_id)) {
                        $intersection_steps = min($intersection_steps, $wire_count[$wire_id]);
                    }
                    [$debug_grid, $intersections] = drawCharacter($debug_grid, $xpos, $ypos, $direction, $wire_id, $intersections);
                }
                break;
            case DOWN:
                for ($i = 0; $i < $distance; $i++) {
                    $xpos--;
                    $wire_count[$wire_id]++;
                    if (hasIntersection($debug_grid, $xpos, $ypos, $wire_id)) {
                        $intersection_steps = min($intersection_steps, $wire_count[$wire_id]);
                    }
                    [$debug_grid, $intersections] = drawCharacter($debug_grid, $xpos, $ypos, $direction, $wire_id, $intersections);
                }
                break;
            case RIGHT:
                for ($i = 0; $i < $distance; $i++) {
                    $ypos++;
                    $wire_count[$wire_id]++;
                    if (hasIntersection($debug_grid, $xpos, $ypos, $wire_id)) {
                        $intersection_steps = min($intersection_steps, $wire_count[$wire_id]);
                    }
                    [$debug_grid, $intersections] = drawCharacter($debug_grid, $xpos, $ypos, $direction, $wire_id, $intersections);
                }
                break;
        }
        // We will only be here after all the movements are done, the next one will (probaly?) be a turn so display a '+'
//        $debug_grid[$xpos][$ypos] = "+";
    }
}

//debug($debug_grid);
echo PHP_EOL;
echo "Part 1: " . findPart1($intersections) . PHP_EOL;
echo "Part 2: " . $intersection_steps . PHP_EOL;

function hasIntersection($grid, $x, $y, $wire_id): bool
{
    if (isset($grid[$x][$y]) && $grid[$x][$y] !== $wire_id) {
        return true;
    }

    return false;

//    return isset($grid[$x][$y]) && $wire_id !== $grid[$x][$y];
}

function drawCharacter($grid, $x, $y, $direction, $wire_id, $intersections): array
{
    if (hasIntersection($grid, $x, $y, $wire_id)) {
        $intersections[] = [$x, $y];
    }
    switch ($direction) {
        case UP:
        case DOWN:
//            $grid[$x][$y] = "|";
            $grid[$x][$y] = $wire_id;
            break;
        case LEFT:
        case RIGHT:
//            $grid[$x][$y] = "-";
            $grid[$x][$y] = $wire_id;
            break;
    }
    return [$grid, $intersections];
}

function findPart1($grid)
{
    $lowest_distance = PHP_INT_MAX;

    foreach ($grid as [$x, $y]) {
        $lowest = abs($x) + abs($y);
        $lowest_distance = min($lowest_distance, $lowest);
    }

//    foreach ($grid as $x => $ys) {
//        foreach ($ys as $y => $cell) {
//            if ($cell === "X") {
//                $lowest = abs($x) + abs($y);
//                $lowest_distance = min($lowest_distance, $lowest);
//            }
//        }
//    }

    return $lowest_distance;
}

function debug($grid)
{
    $xamount = 50;
    $yamount = 150;

    for ($x = $xamount; $x >= -2; $x--) {
        for ($y = -2; $y < $yamount; $y++) {
            if (!isset($grid[$x][$y])) {
                echo ".";
            } else {
                echo $grid[$x][$y];
            }
        }
        echo PHP_EOL;
    }
}
