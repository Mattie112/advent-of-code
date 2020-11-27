<?php
/**
 * https://adventofcode.com/2018/day/11
 */


$serial_number = 4455;

function getPowerLevel($serial_number, $xreq, $yreq)
{
    $grid = generateGrid($serial_number);

    return $grid[$xreq][$yreq];
}

echo getPowerLevel(8, 3, 5) . PHP_EOL;
echo getPowerLevel(57, 122, 79) . PHP_EOL;
echo getPowerLevel(39, 217, 196) . PHP_EOL;
echo getPowerLevel(71, 101, 153) . PHP_EOL;

$maxdata = findSquare($serial_number, 300, 3);
echo ("Part #1 max: " . $maxdata[0]) . " at " . $maxdata[1][0] . "," . $maxdata[1][1] . PHP_EOL;

$part2 = null;
$maxsize = 0;
$grid = generateGrid($serial_number, 300);

for ($i = 1; $i < 300; $i++) {
    echo "Searching for size " . $i . PHP_EOL;
    $maxdata = findSquare($serial_number, 300, $i, $grid);
    if ($part2 === null || $maxdata[0] > $part2[0]) {
        echo ("Found more powerfull grid for size " . $i . " - max: " . $maxdata[0]) . " at " . $maxdata[1][0] . "," . $maxdata[1][1] . PHP_EOL;
        $part2 = $maxdata;
        $maxsize = $i;
    }
}
echo ("Part #2 max: " . $part2[0]) . " at " . $part2[1][0] . "," . $part2[1][1] . "," . $maxsize . PHP_EOL;


function findSquare($serial_number, $size = 300, $search_size = 3, $grid = null)
{
    if ($grid === null) {
        $grid = generateGrid($serial_number, $size);
    }

    $max_power = -999999999;
    $max_coordinates = [];
    foreach ($grid as $y => $row) {
        foreach ($row as $x => $power_level) {
            $power = 0;
            for ($i = 0; $i < $search_size; $i++) {
                for ($j = 0; $j < $search_size; $j++) {
                    if (isset($grid[$x + $i][$y + $j])) {
                        $power += $grid[$x + $i][$y + $j];
                    }
                }
            }
            if ($power > $max_power) {
                $max_power = $power;
                $max_coordinates = [$x, $y];
            }
        }
    }

    return [$max_power, $max_coordinates];
}

function generateGrid($serial_number, $size = 300)
{
    $grid = [];
    for ($y = 1; $y <= $size; $y++) {
        for ($x = 1; $x <= $size; $x++) {
            $rack_id = $x + 10;
            $power_level = $rack_id * $y;
            $power_level += $serial_number;
            $power_level *= $rack_id;
            if ($power_level >= 100) {
                $power_level = substr($power_level, -3, 1);
            } else {
                $power_level = 0;
            }
            $power_level -= 5;
            $grid[$x][$y] = $power_level;
        }
    }

    return $grid;
}

