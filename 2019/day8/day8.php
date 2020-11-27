<?php
/** https://adventofcode.com/2019/day/8 */

$input = trim(file_get_contents(__DIR__ . "/day8-input.txt"));
$width = 25;
$height = 6;

$layers = [];
$layer_id = 1;
while (!empty($input)) {
    $layers[$layer_id] = [];
    for ($i = 0; $i < $height; $i++) {
        $layers[$layer_id][] = substr($input, $i * $width, $width);
    }
    $input = substr($input, $width * $height);
    $layer_id++;
}

// Find the layer with the fewest 0 digits
$lowest_zero_digit_count = PHP_INT_MAX;
$lowest_zero_layer_id = null;
foreach ($layers as $layer_id => $layer) {
    $zero_digit_count = 0;
    foreach ($layer as $row) {
        $chars = count_chars($row, 1);
        $zero_digit_count += $chars[ord("0")];
    }

    if ($zero_digit_count < $lowest_zero_digit_count) {
        $lowest_zero_digit_count = $zero_digit_count;
        $lowest_zero_layer_id = $layer_id;
    }
}

echo "Layer " . $lowest_zero_layer_id . " has the lowest amount of 0 digits (" . $lowest_zero_digit_count . ")" . PHP_EOL;

// Now get the amount of 1's and 2's of that layer
$tmp_layers = $layers[$lowest_zero_layer_id];
$ones = 0;
$twos = 0;
foreach ($tmp_layers as $row) {
    $chars = count_chars($row, 1);
    $ones += $chars[ord("1")];
    $twos += $chars[ord("2")];
}

echo "Part 1: " . ($ones * $twos) . PHP_EOL;

// Now on to part 2: decoding the image
$image = [];
// Loop over layers (in reverse so we know what to overwrite)
$layers = array_reverse($layers);
foreach ($layers as $layer_id => $layer) {
    foreach ($layer as $row_id => $row) {
        $row = str_split($row);
        foreach ($row as $pixel_id => $pixel) {
            if ($pixel !== "2") {
                $image[$row_id][$pixel_id] = $pixel;
            }
        }
    }
}

echo "Part 2:" . PHP_EOL;
for ($i = 0; $i < $height; $i++) {
    for ($y = 0; $y < $width; $y++) {
        if ($image[$i][$y] === "1") {
            echo "#";
        } else {
            echo " ";
        }
    }
    echo PHP_EOL;
}
