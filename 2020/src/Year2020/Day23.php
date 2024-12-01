<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day23 extends Day
{
    public int $moves = 100;

    public function part1(): int|string
    {
        $input = $this->getInput(2020, 23, 1);
        $input = str_split($input);
        $circle = array_map("intval", $input);

        $circle_count = count($circle);

        $move = 1;
        $current_cup_id = 0;
        while ($move <= $this->moves) {
            $destination_cup = false;
            $destination_cup_id = false;
            $this->log(sprintf("-- move %d --", $move));
            $current_cup = $circle[$current_cup_id];

            if ($this->output?->isVerbose()) {
                foreach ($circle as $id => $cup) {
                    if ($id === $current_cup_id) {
                        $this->logNoLineBreak("(" . $cup . ")");
                    } else {
                        $this->logNoLineBreak($cup);
                    }
                    $this->logNoLineBreak(" ");
                }
                $this->log("");
            }

            $cup1 = $circle[($current_cup_id + 1) % $circle_count];
            $cup2 = $circle[($current_cup_id + 2) % $circle_count];
            $cup3 = $circle[($current_cup_id + 3) % $circle_count];
            $this->log(sprintf("pick up: %d, %d, %d", $cup1, $cup2, $cup3));
            unset($circle[($current_cup_id + 1) % $circle_count], $circle[($current_cup_id + 2) % $circle_count], $circle[($current_cup_id + 3) % $circle_count]);
            $circle = array_values($circle); // Reset our keys so that I can use the length in my array_slices

            $tmp = $current_cup - 1;
            while ($destination_cup === false) {
                foreach ($circle as $id => $t) {
                    if ($t === $tmp) {
                        $destination_cup = $t;
                        $destination_cup_id = $id;
                        break 2;
                    }
                }
                $tmp--;
                if ($tmp < min($circle)) {
                    $tmp = max($circle);
                }
            }
            $this->log(sprintf("destination: %d", $destination_cup));

            $first_part = array_slice($circle, 0, $destination_cup_id + 1, false);
            $second_part = array_slice($circle, $destination_cup_id + 1, count($circle), false);
            $circle = array_merge($first_part, [$cup1, $cup2, $cup3], $second_part);
            $circle_count = count($circle);

            // The structure of our circle might be changed (well because it is a circle this shouldn't be but hey I'm just using an array) so we need to determine the ID of our current cup again
            $current_cup_id = array_search($current_cup, $circle, false);

            // Now determine the next one and move on!
            $current_cup_id = (($current_cup_id + 1) % $circle_count);
            $move++;
            $this->log("");
        }

        $this->log(implode(", ", $circle));

        // Now find cup 1
        $current_cup_id = array_search(1, $circle, false);
        $answer = "";
        while (count($circle) > 0) {
            $current_cup_id = ($current_cup_id + 1) % $circle_count;
            $answer .= $circle[$current_cup_id];
            unset($circle[$current_cup_id]);
        }
        $answer = substr($answer, 0, -1);

        return $answer;

    }

    public function part2(): int|string
    {
        ini_set("memory_limit", "-1");
        return 0;
        $input = $this->getInput(2020, 23, 2);
        $input = str_split($input);
        $circle = array_map("intval", $input);

        $a = range(max($circle), 1000000);
        $circle = array_merge($circle, $a);

        $circle_count = count($circle);

        $move = 1;
        $current_cup_id = 0;
        while ($move <= 10000000) {
            $destination_cup = false;
            $destination_cup_id = false;
            $this->log(sprintf("-- move %d --", $move));
            $current_cup = $circle[$current_cup_id];

            if ($this->output?->isVerbose()) {
                foreach ($circle as $id => $cup) {
                    if ($id === $current_cup_id) {
                        $this->logNoLineBreak("(" . $cup . ")");
                    } else {
                        $this->logNoLineBreak($cup);
                    }
                    $this->logNoLineBreak(" ");
                }
                $this->log("");
            }

            $cup1 = $circle[($current_cup_id + 1) % $circle_count];
            $cup2 = $circle[($current_cup_id + 2) % $circle_count];
            $cup3 = $circle[($current_cup_id + 3) % $circle_count];
            $this->log(sprintf("pick up: %d, %d, %d", $cup1, $cup2, $cup3));
            unset($circle[($current_cup_id + 1) % $circle_count], $circle[($current_cup_id + 2) % $circle_count], $circle[($current_cup_id + 3) % $circle_count]);
            $circle = array_values($circle); // Reset our keys so that I can use the length in my array_slices

            $tmp = $current_cup - 1;
            while ($destination_cup === false) {
                foreach ($circle as $id => $t) {
                    if ($t === $tmp) {
                        $destination_cup = $t;
                        $destination_cup_id = $id;
                        break 2;
                    }
                }
                $tmp--;
                if ($tmp < min($circle)) {
                    $tmp = max($circle);
                }
            }
            $this->log(sprintf("destination: %d", $destination_cup));

            $first_part = array_slice($circle, 0, $destination_cup_id + 1, false);
            $second_part = array_slice($circle, $destination_cup_id + 1, count($circle), false);
            $circle = array_merge($first_part, [$cup1, $cup2, $cup3], $second_part);
            $circle_count = count($circle);

            // The structure of our circle might be changed (well because it is a circle this shouldn't be but hey I'm just using an array) so we need to determine the ID of our current cup again
            $current_cup_id = array_search($current_cup, $circle, false);

            // Now determine the next one and move on!
            $current_cup_id = (($current_cup_id + 1) % $circle_count);
            $move++;
            $this->log("");
        }

        $this->log(implode(", ", $circle));

        // Now find cup 1
        $current_cup_id = array_search(1, $circle, false);
        $answer = "";
        while (count($circle) > 0) {
            $current_cup_id = ($current_cup_id + 1) % $circle_count;
            $answer .= $circle[$current_cup_id];
            unset($circle[$current_cup_id]);
        }
        $answer = substr($answer, 0, -1);

        return $answer;
    }


}
