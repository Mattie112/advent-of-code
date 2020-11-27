<?php
/**
 * http://adventofcode.com/2017/day/6
 */

$answer = 0;

if ($file = fopen(__DIR__ . "/day6-input.txt", "rb")) {
    while (!feof($file)) {
        $banks = trim(fgets($file));

        if (empty($banks)) {
            continue;
        }

        // Accept tabs and space
        $banks = preg_split('/\s+/', $banks);
        $banks_amount = count($banks);
        $seen_states = [];

        $go = true;
        while ($go) {
            $most_blocks = array_search(max($banks), $banks, false);
            $blocks_to_devide = $banks[$most_blocks];
            $banks[$most_blocks] = 0;
            $i = $most_blocks + 1;
            $loop_array = true;

            while ($blocks_to_devide > 0) {
                if ($i > $banks_amount - 1) {
                    $i = 0;
                }
                $bank = &$banks[$i];
                $bank++;
                $blocks_to_devide--;
                $i++;
            }

            $state = implode(" ", $banks);
            $answer++;
            if (in_array($state, $seen_states, false)) {
                echo $answer . PHP_EOL;
                die();

            }
            $seen_states[] = $state;
        }
    }
    fclose($file);
}


