<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day13 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 13, 1);
        $my_timestamp = (int)$input[0];
        $busses = array_map("intval", array_filter(explode(",", $input[1]), static fn($elem) => $elem !== "x"));

        $timestamp = $my_timestamp;

        while (true) {
            foreach ($busses as $bus) {
                if ($timestamp % $bus === 0) {
                    // A bus can leave at this time and because of the return we also know our answer here :)
                    return ($timestamp - $my_timestamp) * $bus;
                }
            }
            $timestamp++;
        }
    }

    // https://www.reddit.com/r/adventofcode/comments/kc60ri/2020_day_13_can_anyone_give_me_a_hint_for_part_2/
    // https://www.reddit.com/r/adventofcode/comments/kc4njx/2020_day_13_solutions/gfnztpp/?utm_source=reddit&utm_medium=web2x&context=3
    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 13, 2);
        $busses = array_map("intval", array_filter(explode(",", $input[1]), static fn($elem) => $elem !== "x"));

        $multiplier = $busses[0]; // So 7 with our test input
        unset($busses[0]);
        $i = 0;
        foreach ($busses as $bus_key => $bus) {
            while (true) {
                if (($i + $bus_key) % $bus === 0) {
                    $multiplier *= $bus;
                    break;
                }
                $i += $multiplier;
            }
        }

        return $i;
    }
}
