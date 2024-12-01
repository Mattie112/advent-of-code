<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use JetBrains\PhpStorm\Pure;
use mattie112\AdventOfCode\Day;

class Day9 extends Day
{
    protected int $preamble = 25;

    public function setUp($input): array
    {
        $this->preamble = 25;
        if ($this->isTest()) {
            $this->preamble = 5;
        }
        $input = array_map('intval', $input);
        return $input;
    }

    public function part1(): int|string
    {
        $input = $this->setUp($this->getInputAsArray(2020, 9, 1));
        $count = count($input);

        // Now go and check the rest
        for ($i = $this->preamble; $i < $count; $i++) {
            $array_slice = array_slice($input, $i - ($this->preamble), $this->preamble);
            $valid = $this->isValid($array_slice, $input[$i]);
            if (!$valid) {
                return $input[$i];
            }
        }

        return 0;
    }

    public function isValid($arr, $number): bool
    {
        foreach ($arr as $a) {
            foreach ($arr as $b) {
                if ($a + $b === $number && $a !== $b) {
                    return true;
                }
            }
        }
        return false;
    }

    public function part2(): int|string
    {
        $input = $this->setUp($this->getInputAsArray(2020, 9, 2));
        $count = count($input);

        // This is the same code as part #1 however we don't return value here but we start looking with that value for a continuous  set
        for ($i = $this->preamble; $i < $count; $i++) {
            $array_slice = array_slice($input, $i - ($this->preamble), $this->preamble);
            $valid = $this->isValid($array_slice, $input[$i]);
            if (!$valid) {
                return $this->findContinuousSet($input, $input[$i]);
            }
        }

        return 0;
    }

    #[Pure] public function findContinuousSet($arr, $wanted_number): int
    {
        foreach ($arr as $i => $value) {
            if ($value >= $wanted_number) {
                // We can ignore these anyway
                continue;
            }

            // Now lets find some stuff
            $counter = $min = $max = $value;
            // We ignore all previous values
            $j = $i;
            // As soon as our counter exeeds the number that we want we can halt execution
            while ($counter < $wanted_number) {
                $j++;
                $counter += $arr[$j];
                $min = min($min, $arr[$j]);
                $max = max($max, $arr[$j]);
            }

            // If we do have the number that we want we can return the $min + $max
            if ($counter === $wanted_number) {
                return ($min + $max);
            }
        }
        return 0;
    }
}
