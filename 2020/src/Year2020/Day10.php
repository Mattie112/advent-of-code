<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day10 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 10, 1);
        sort($input);
        $device_rating = max($input) + 3;
        $input[] = $device_rating;
        $input = array_map('intval', $input);

        $prev = 0;
        $diff_1 = [];
        $diff_3 = [];
        foreach ($input as $jolt) {
            $diff = $jolt - $prev;

            if ($diff === 1) {
                $diff_1[] = $jolt;
            }

            if ($diff === 3) {
                $diff_3[] = $jolt;
            }

            $prev = $jolt;
        }

        return (count($diff_1) * count($diff_3));
    }

    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 10, 2);
        sort($input);
        $device_rating = max($input) + 3;
        $input[] = $device_rating;
        $input = array_map('intval', $input);

        $adapters[0] = 1;
        foreach ($input as $id => $jolt) {
            $sum = 0;
            for ($i = $jolt - 3; $i < $jolt; $i++) {
                if (isset($adapters[$i])) {
                    $sum += $adapters[$i];
                }
            }
            $adapters[$jolt] = $sum;
        }
        return $adapters[max($input)];
    }
}
