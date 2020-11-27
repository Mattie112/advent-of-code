<?php
/**
 * https://adventofcode.com/2018/day/3
 */

$fabric = [];
$duplicate_fabric = 0;
$claims = [];
$single_claim = null;

if ($file = fopen(__DIR__ . "/day3-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if (preg_match("/#(\d+) @ (\d+),(\d+): (\d+)x(\d+)/", $line, $matches)) {
            $ints = array_map('intval', $matches);
            $claims[] = array_map('intval', $ints);
            list(, $id, $from_left, $from_top, $width, $height) = array_map('intval', $ints);

            for ($x = $from_top; $x < $from_top + $height; $x++) {
                for ($y = $from_left; $y < $from_left + $width; $y++) {
                    // Only count duplicate fabric once
                    if (isset($fabric[$y][$x]) && $fabric[$y][$x] === 1) {
                        $duplicate_fabric++;
                    }

                    if (!isset($fabric[$y][$x])) {
                        $fabric[$y][$x] = 1;
                    } else {
                        $fabric[$y][$x]++;
                    }
                }
            }
        }
    }
    fclose($file);
}

// Now check for claims that have no overlap by looping over the claims again and for each claim
// check if the ENTIRE $x,$y only has a single claim, if that is the case stop the loop because we found our answer!
foreach ($claims as list(, $id, $from_left, $from_top, $width, $height)) {
    $x_ok = 0;
    for ($x = $from_top; $x < $from_top + $height; $x++) {
        $y_ok = false;
        $y_count = 0;
        for ($y = $from_left; $y < $from_left + $width; $y++) {
            if ($fabric[$y][$x] === 1) {
                $y_count++;

                // Ok our $y matches our $width we should now go to the next $x
                if ($y_count === $width) {
                    $y_ok = true;
                }
                continue;
            }
            $y_count = 0;
        }
        // If your $y was OK, mark this $x as OK
        if ($y_ok) {
            $x_ok++;
        } else {
            // If not, reset it
            $x_ok = 0;
        }
        // If our $x ok matches our hiehgt we are done!
        if ($x_ok === $height) {
            $single_claim = $id;
            break 2;
        }
    }
}

echo ("Part #1: " . $duplicate_fabric . " inches in >= 2 claims") . PHP_EOL;
echo ("Part #2: #" . $single_claim . " is not duplicated by anyone") . PHP_EOL;
