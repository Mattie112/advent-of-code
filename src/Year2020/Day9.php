<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use JetBrains\PhpStorm\Pure;
use mattie112\AdventOfCode\Day;

class Day9 extends Day
{
    public function part1(): int|string
    {
        $preamble = 25;
        if ($this->isTest()) {
            $preamble = 5;
        }
        $input = $this->getInputAsArray(2020, 9, 1);
        $input = array_map('intval', $input);
        $count = count($input);

        // Now go and check the rest
        for ($i = $preamble; $i < $count; $i++) {
            $array_slice = array_slice($input, $i - ($preamble), $preamble);
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
        $preamble = 25;
        if ($this->isTest()) {
            $preamble = 5;
        }
        $input = $this->getInputAsArray(2020, 9, 2);
        $input = array_map('intval', $input);
        $count = count($input);

        // Now go and check the rest
        for ($i = $preamble; $i < $count; $i++) {
            $array_slice = array_slice($input, $i - ($preamble), $preamble);
            $valid = $this->isValid($array_slice, $input[$i]);
            if (!$valid) {
                return $this->findContiniousSet($input, $input[$i]);
            }
        }

        return 0;
    }

    #[Pure] public function findContiniousSet($arr, $number)
    {
        $count = count($arr);
        for ($i = 0; $i < $count; $i++) {
            if ($arr[$i] >= $number) {
                continue;
            }

            // Now lets find some stuff
            $counter = $arr[$i];
            $min = $counter;
            $max = $counter;
            $j = $i;
            while ($counter < $number) {
                $j++;
                $tmp = $arr[$j];
                $counter += $tmp;
                $min = min($min, $arr[$j]);
                $max = max($max, $arr[$j]);
            }

            if ($counter === $number) {
                return ($min + $max);
            }
        }
        return "bleh";
    }


}
