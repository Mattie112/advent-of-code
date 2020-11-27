<?php
/**
 * https://adventofcode.com/2018/day/6
 */

$points = [];
$max_x = PHP_INT_MIN;
$max_y = PHP_INT_MIN;
$min_x = PHP_INT_MAX;
$min_y = PHP_INT_MAX;
if ($file = fopen(__DIR__ . "/day6-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if (preg_match("@(\d+), (\d+)@", $line, $matches)) {
            $x = (int) $matches[1];
            $y = (int) $matches[2];
            $points[] = [$x, $y];
            $max_x = max($x, $max_x);
            $max_y = max($y, $max_y);
            $min_x = min($x, $min_x);
            $min_y = min($y, $min_y);
        }
    }
    fclose($file);
}

// Now I know the size of my grid:
echo "Lop left: x:" . $min_x . " y:" . $min_y . PHP_EOL;
echo "Bottom rigt: x:" . $max_x . " y:" . $max_y . PHP_EOL;

$distances = [];
$closest_points = [];
$sums = [];

// We will create the entire grid (by min/max x/y instead of infinite)
// For each x/y we will calculate the distance to all points
// Is we are lower then whatever we had it means we are closer so we can overwrite it
// Ties are set to -1 as nothing can be lower then that
for ($y = $min_y; $y < $max_y; $y++) {
    for ($x = $min_x; $x < $max_x; $x++) {
        // Init this grid cell
        $distances[$y][$x] = PHP_INT_MAX;
        $closest_points[$y][$x] = -1;
        $sums[$y][$x] = 0;

        // Now for each point calculate the distance
        foreach ($points as $point_id => list($point_x, $point_y)) {
            $manhattan = abs($y - $point_x) + abs($x - $point_y);
            $sums[$y][$x] += $manhattan;

            if ($manhattan < $distances[$y][$x]) {
                $distances[$y][$x] = $manhattan;
                $closest_points[$y][$x] = $point_id;
            } elseif ($manhattan === $distances[$y][$x]) {
                $closest_points[$y][$x] = -1; // A tie
            }
        }
    }
}
// Not sure why but it seems that looping X -> Y (and $var[x][y]) gives a different result then Y -> X (and $var[y][x])

// Now claculate the size of each 'closest' area
$areas = [];
for ($x = $min_x; $x < $max_x; $x++) {
    for ($y = $min_y; $y < $max_y; $y++) {
        $point_id = $closest_points[$x][$y];
        $areas[$point_id]++;
    }
}
// We don't care about ties
unset($areas[-1]);
// Now we have an array with [point_id => area_size] but this includes "infinite" areas

// All areas that are at the edge should be removed as the are infinite
for ($x = $min_x; $x <= $max_x; $x++) {
    $point_id = $closest_points[$x][$min_y];
    unset($areas[$point_id]);
    $point_id = $closest_points[$x][$max_y];
    unset($areas[$point_id]);
}
for ($y = $min_y; $y <= $max_y; $y++) {
    $point_id = $closest_points[$min_x][$y];
    unset($areas[$point_id]);
    $point_id = $closest_points[$max_x][$y];
    unset($areas[$point_id]);
}

//  So Finally our answer
$tmp = array_keys($areas, max($areas));
echo "The largest non-infinite area = " . max($areas) . " (#" . reset($tmp) . ")" . PHP_EOL;

// Part 2
$total_safe_area = 0;
for ($x = $min_x; $x <= $max_y; $x++) {
    for ($y = $min_y; $y <= $max_y; $y++) {
        $tmp_area = 0;
        foreach ($points as [$point_x, $point_y]) {
            $tmp_area += abs($x - $point_x) + abs($y - $point_y);
        }

        if ($tmp_area < 10000) {
            $total_safe_area++;
        }
    }
}

echo "Total safe area " . $total_safe_area . PHP_EOL;
