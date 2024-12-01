<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day15 extends Day
{
    public function part1(): int|string
    {
        ini_set("memory_limit", "-1");
        $input = $this->getInputAsArray(2020, 15, 1)[0];

        return $this->solve($input, 2020);
    }


    public function part2(): int|string
    {
        ini_set("memory_limit", "-1");
        $input = $this->getInputAsArray(2020, 15, 2)[0];

        return $this->solve($input, 30000000);
    }

    public function solve(string $input_str, int $dinner_is_ready): int
    {
        $input = explode(",", $input_str);
        $last = array_pop($input);
        $input_count = count($input);
        $memory = array_combine($input, range(1, $input_count)); // Use range so I dont have turn 0
        $turn = $input_count + 1;
        while ($turn < $dinner_is_ready) {
            $turn++;
            $prev_turn = $turn - 1;
            if (!isset($memory[$last])) {
                $memory[$last] = $prev_turn;
                $last = 0;
            } else {
                $tmp_previous_turn = $memory[$last];
                $memory[$last] = $prev_turn;
                $last = $prev_turn - $tmp_previous_turn;
            }
        }

        return $last;
    }
}
