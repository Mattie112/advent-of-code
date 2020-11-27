<?php
/**
 * http://adventofcode.com/2017/day/10
 */

$list = range(0, 255);

// Read the file and store in array
if ($file = fopen(__DIR__ . "/day10-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if ($line === "") {
            continue;
        }

        $inputs = explode(",", $line);
    }
    fclose($file);
}

$list_pos = 0;
$skip_size = 0;
$listcount = count($list);

foreach ($inputs as $input) {
    $sublist = [];
    for ($i = 0; $i < $input; $i++) {
        $sublist[] = $list[($list_pos + $i) % $listcount];
    }

    $sublist = array_reverse($sublist);
    for ($i = 0; $i < $input; $i++) {
        $list[($list_pos + $i) % $listcount] = $sublist[$i];
    }

    $list_pos = $list_pos + $input + $skip_size;
    $skip_size++;
}

echo "Found answer (part1): " .($list[0] * $list[1]).PHP_EOL;
