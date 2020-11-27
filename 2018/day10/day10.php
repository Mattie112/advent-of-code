<?php
/**
 * https://adventofcode.com/2018/day/10
 */

$points = [];

if ($file = fopen(__DIR__ . "/day10-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }

        // Find and store the coordinated & velocities
        if (preg_match("@position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>@", $line, $matches)) {
            $x = (int) $matches[1];
            $y = (int) $matches[2];
            $right = (int) $matches[3];
            $down = (int) $matches[4];

            $points[] = [$x, $y, $right, $down];
            echo "Point moves from " . $x . "," . $y . " with a speed of RIGHT:" . $right . " and DOWN:" . $down . PHP_EOL;
        }
    }
}

// If we start 'gaining' height we know we should not continue as I assume the letters are visible when it is most compact
// we don't really need a prev_sky but as we don't display each step (performance) we do need to keep it
$height = null;
$prev_height = null;
$prev_sky = null;
$steps = 0;
do {
    $prev_height = $height;
    $prev_sky = $points;
    foreach ($points as &$point) {
        $point[0] += $point[2];
        $point[1] += $point[3];
    }
    unset($point);
    $height = max(array_column($points, 1));
    $steps++;
} while ($prev_height === null || $height < $prev_height);

drawSky($prev_sky);

echo PHP_EOL . "DONE (see output above!) We only needed to wait " . ($steps - 1) . " seconds  or " . round(($steps - 1) / 60 / 60, 2) . " hours for it" . PHP_EOL;

// It will look amazing I promise! I even used the star symbol so it looks even more like a sky!
function drawSky($points)
{
    echo PHP_EOL;
    // Find min/max so we know how much to draw
    $maxX = max(array_column($points, 0));
    $minX = min(array_column($points, 0));
    $maxY = max(array_column($points, 1));
    $minY = min(array_column($points, 1));
    echo "We have a sky from minX:" . $minX . " to maxX:" . $maxX . " and minY:" . $minY . " to maxY:" . $maxY . PHP_EOL . PHP_EOL;

    $sky = [];

    // Sort points based on their x/y
    foreach ($points as $point) {
        $sky[$point[0]][$point[1]] = true;
    }

    // Let's draw it!
    for ($y = $minY; $y <= $maxY; $y++) {
        for ($x = $minX; $x <= $maxX; $x++) {
            if (isset($sky[$x][$y])) {
                echo "*";
            } else {
                echo " ";
            }
        }
        echo PHP_EOL;
    }
}

