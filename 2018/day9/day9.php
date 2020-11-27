<?php
/**
 * https://adventofcode.com/2018/day/9
 */

$debug = false;

if ($file = fopen(__DIR__ . "/day9-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }

        $players = [];
        if (preg_match("@(\d+) players.* (\d+) points@", $line, $matches)) {
            $player_amount = (int) $matches[1];
            $players = range(1, $player_amount);
            $last_marble = (int) $matches[2];
            echo "Calculating game for " . $player_amount . " players with " . $last_marble . " marbles" . PHP_EOL;

            $circle = [1 => 0];
            $current_marble = 1;
            $player = 0;
            $scores = array_fill(0, $player_amount, 0);
            if ($debug) {
                echo "[-] (0)" . PHP_EOL;
            }
            $time = 0;
            $last_time = 0;
            $time_between_log = 0;
            for ($i = 1; $i < $last_marble; $i++) {
                $time = microtime(true);
                if ($i % 100 === 0) {
                    echo $i . " / " . $last_marble . " (" . $last_time . " seconds for last marble), Time between log (100 marbles): " . (microtime(true) - $time_between_log) . PHP_EOL;
                    $time_between_log = microtime(true);
                }
                $player = (($player + 1) % $player_amount);
                if ($debug) {
                    echo "[" . $player . "] ";
                }
                $circle_count = count($circle);
                if ($i % 23 === 0) {
                    $remove_marble_i = (($circle_count + $current_marble - 7) % $circle_count + 1);
                    $removed_marble = $circle[$remove_marble_i];
                    array_splice($circle, $remove_marble_i, 1);
                    $scores[$player] += $i;
                    $scores[$player] += $removed_marble;
                    $current_marble = $remove_marble_i - 1;
                } else {
                    $new_pos = ($current_marble + 2) % $circle_count;
                    array_splice($circle, $new_pos + 1, 0, $i);
                    $current_marble = $new_pos;
                }

                // DEBUG STUFF
                if ($debug) {
                    for ($j = 0, $jMax = count($circle); $j < $jMax; $j++) {
                        if ($j === $current_marble) {
                            echo "(" . $circle[$j] . ") ";
                        } else {
                            echo $circle[$j] . " ";
                        }
                    }
                    echo PHP_EOL;
                }
                $last_time = microtime(true) - $time;
            }

            echo max($scores) . PHP_EOL;
        }
    }
}

