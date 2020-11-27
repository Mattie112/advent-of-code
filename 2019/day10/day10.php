<?php
/** https://adventofcode.com/2019/day/10 */

$map = [];
if ($file = fopen(__DIR__ . "/day10-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }
        $map[] = str_split($line);
    }
}
$astroid_map = [];
$most_astroids = 0;
$best_loc_x = 0;
$best_loc_y = 0;
foreach ($map as $row_id => $row) {
    foreach ($row as $cell_id => $cell) {
        if ($cell === ".") {
            continue;
        }
        $seen = [];
        foreach ($map as $yy => $rowY) {
            foreach ($rowY as $xx => $cellX) {
                if ($map[$yy][$xx] === "#" && ($yy !== $row_id || $xx !== $cell_id)) {
                    $degrees = rad2deg(atan2($yy - $row_id, $xx - $cell_id));
                    $seen[] = $degrees;
                }
            }
        }
        $count = count(array_unique($seen));
        if ($count > $most_astroids) {
            $most_astroids = $count;
            $best_loc_x = $cell_id;
            $best_loc_y = $row_id;
        }
    }
}
echo "Part 1: " . $best_loc_x . "," . $best_loc_y . " (" . $most_astroids . " astroids)" . PHP_EOL;

// Now that we have our best position go and calculate the astroid line of sight for this position (same as the function above but simply run again for a single position)
$astroid_list = getVisibleAstroidsForPosition($map, $best_loc_y, $best_loc_x);

ksort($astroid_list);

// Now we can loop (clockwise, degrees INCREASING) through our astroids and simple kill the closest we can find
$kill_count = 0;
while (true) {
    foreach ($astroid_list as $degrees => $astroids) {
        if (empty($astroids)) {
            continue;
        }
        $key = min(array_keys($astroids));
        $closest = $astroids[$key];
        unset($astroid_list[$degrees][$key]);
        $kill_count++;
        echo $kill_count . " - " . $degrees . " - " . implode(",", $closest) . PHP_EOL;
        if ($kill_count === 200) {
            $part2 = ($closest[0] * 100) + $closest[1];
            echo "Part 2: " . $part2 . PHP_EOL;
            break 2;
        }
    }
}

function getVisibleAstroidsForPosition($map, $best_loc_y, $best_loc_x)
{
    $astroid_list = [];
    foreach ($map as $row_id => $row) {
        foreach ($row as $cell_id => $cell) {
            if ($cell === "." || ($row_id === $best_loc_y && $cell_id === $best_loc_x)) {
                continue;
            }
            $degrees = getDegrees($best_loc_y - $row_id, $best_loc_x - $cell_id);
            $astroid_list[(string)$degrees][abs($best_loc_y - $row_id) + abs($best_loc_x - $cell_id)] = [$cell_id, $row_id];
        }
    }
    return $astroid_list;
}

function getDegrees($dX, $dY)
{
    if ($dX === 0 && $dY === 0) {
        return 0;
    }
    $degrees = rad2deg(atan2($dX, $dY));
    $degrees -= 90.0;
    // All negative values will get +360 degrees so that our clockwise layer can just go through the array from low to high
    if ($degrees < 0) {
        $degrees += 360;
    }
    return $degrees;
}


