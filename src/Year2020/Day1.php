<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;
use RuntimeException;

class Day1 extends Day
{
    public function part1(): int|string
    {
        $input = explode("\n", $this->getInput(2020, 1, 1));

        foreach ($input as $i) {
            $i = (int)$i;
            foreach ($input as $j) {
                $j = (int)$j;
                if ($i + $j === 2020) {
                    return ($i * $j);
                }
            }
        }
        throw new RuntimeException("Did not find an answer");
    }

    public function part2(): int|string
    {
        $input = explode("\n", $this->getInput(2020, 1, 2));

        foreach ($input as $i) {
            $i = (int)$i;
            foreach ($input as $j) {
                $j = (int)$j;
                foreach ($input as $k) {
                    $k = (int)$k;
                    if ($i + $j + $k === 2020) {
                        return ($i * $j * $k);
                    }
                }
            }
        }
        throw new RuntimeException("Did not find an answer");
    }
}
