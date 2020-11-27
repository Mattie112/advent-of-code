<?php

/**
 * https://adventofcode.com/2018/day/13
 */

const HORIZONTAL = "-";
const VERTICAL = "|";
const RIGHT_TO_LEFT = "/";
const LEFT_TO_RIGHT = "\\";
const INTERSECTION = "+";
const CART_UP = "^";
const CART_DOWN = "v";
const CART_LEFT = "<";
const CART_RIGHT = ">";
const DIRECTIONS = [CART_UP, CART_RIGHT, CART_DOWN, CART_LEFT];

$tracks = [];
$carts = [];
$carts_by_yx = [];

if ($file = fopen(__DIR__ . "/day13-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file), "\n\r");
        if ($line === "") {
            continue;
        }

        // Yes this is now y,x as we need to loop line by line lateron
        $tracks[] = str_split($line);

        // Find carts and replace them by the correct tracks
        foreach ($tracks as $y => $track_row) {
            foreach ($track_row as $x => $track) {
                if (in_array($track, [CART_UP, CART_DOWN, CART_LEFT, CART_RIGHT,], true)) {
                    $carts[] = ["y" => $y, "x" => $x, "direction" => $track, "turns" => 0,];
                    $carts_by_yx[$y][$x] = $track;
                    echo "Found cart at line " . $y . " column " . $x . " with position: " . $track . PHP_EOL;
                    if (in_array($track, [CART_UP, CART_DOWN], true)) {
                        $tracks[$y][$x] = VERTICAL;
                    } else {
                        $tracks[$y][$x] = HORIZONTAL;
                    }
                }
            }
        }
    }
}

// Now that we have both tracks & carts we can execute the movement ofthe carts!!!yeah!
$part1 = "";
$iterations = 999999;
for ($a = 0; $a < $iterations; $a++) {
    $carts_to_be_removed = [];
    $width = count($tracks[0]);
    uasort($carts, function ($a, $b) use ($width) {
        return $a["x"] + ($a["y"] * $width) <=> $b["x"] + ($b["y"] * $width);
    });

    foreach ($carts as $i => &$cart) {
        $y = &$cart["y"];
        $x = &$cart["x"];
        $direction = &$cart["direction"];
        $turns = &$cart["turns"];
        echo $x . "x" . $y . PHP_EOL;
    }

    foreach ($carts as $i => &$cart) {
        echo "starting cart " . $i . PHP_EOL;

        $y = &$cart["y"];
        $x = &$cart["x"];
        $direction = &$cart["direction"];
        $turns = &$cart["turns"];

        switch ($direction) {
            case CART_UP:
                $y--;
                break;
            case CART_DOWN;
                $y++;
                break;
            case CART_LEFT;
                $x--;
                break;
            case CART_RIGHT:
                $x++;
                break;
        }

        // Now that that is done see if we need to make a turn
        $track = $tracks[$y][$x];
        switch ($track) {
            // "/"
            case RIGHT_TO_LEFT:
                {
                    switch ($direction) {
                        case CART_UP:
                            $direction = CART_RIGHT;
                            break;
                        case  CART_DOWN:
                            $direction = CART_LEFT;
                            break;
                        case CART_LEFT:
                            $direction = CART_DOWN;
                            break;
                        case CART_RIGHT:
                            $direction = CART_UP;
                    }
                    break;
                }
            // "\"
            case LEFT_TO_RIGHT:
                {
                    switch ($direction) {
                        case CART_UP:
                            $direction = CART_LEFT;
                            break;
                        case  CART_DOWN:
                            $direction = CART_RIGHT;
                            break;
                        case CART_LEFT:
                            $direction = CART_UP;
                            break;
                        case CART_RIGHT:
                            $direction = CART_DOWN;
                            break;
                    }
                    break;
                }
            case INTERSECTION:
                {
                    // We can either to left, strait or right
                    $dir = ($turns % 3) - 1;
                    $current_direction_id = array_search($direction, DIRECTIONS, true);
                    $direction = DIRECTIONS[(4 + $dir + $current_direction_id) % 4];
                    $turns++;

                    break;
                }
        }

        // Now find if we have a crash as that is the most important
        foreach ($carts as $check_index => $cart_to_check) {
            if ($i === $check_index) {
                continue;
            }
            if ($y === $cart_to_check["y"] && $x === $cart_to_check["x"]) {
                debug($carts, $tracks);
                if ($part1 === "") {
                    $part1 = "COLLISSION #1 AT x: " . $x . " - y: " . $y . PHP_EOL;
                }
                echo "COLLISSION AT x: " . $x . " - y: " . $y . PHP_EOL;
                $carts_to_be_removed[] = $check_index;
                $carts_to_be_removed[] = $i;

                echo "Carts left: " . count($carts) . PHP_EOL;
            }
        }
        echo "done with cart " . $i . PHP_EOL;
    }
    foreach ($carts_to_be_removed as $id) {
        unset($carts[$id]);
    }
    if (count($carts) === 1) {
        $tmparr = $carts;
        $tmpcart = reset($tmparr);
        debug($carts, $tracks);
        echo $part1;
        echo "Last cart position x:" . $tmpcart["x"] . " - y:" . $tmpcart["y"] . " after " . $a . " ticks" . PHP_EOL;
        die();
    }
    echo "-------------------------------------" . PHP_EOL;
}

function debug($carts, $tracks)
{
    $carts_by_yx = [];
    foreach ($carts as $tmp) {
        $carts_by_yx[$tmp["y"]][$tmp["x"]] = $tmp["direction"];
    }
    // DEBUG
    for ($j = 0; $j < 200; $j++) {
        for ($k = 0; $k < 200; $k++) {
            if (isset($carts_by_yx[$j][$k])) {
                echo $carts_by_yx[$j][$k];
            } else {
                if (isset($tracks[$j][$k])) {
                    echo $tracks[$j][$k];
                }
            }
        }
        if (isset($tracks[$j])) {
            echo PHP_EOL;
        }
    }
}
