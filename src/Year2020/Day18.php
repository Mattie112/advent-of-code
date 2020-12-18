<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day18 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 18, 1);

        $sum = [];
        foreach ($input as $line) {
            $sum[] = $this->solvePart1($line);
        }

        return array_sum($sum);
    }

    public function solvePart1(string $string)
    {
        $string = trim($string);

        // Recursive thingy to solve every thing between ( ... )
        while (str_contains($string, "(")) {
            if (preg_match("/\(([0-9+* ]+)\)/", $string, $matches)) {
                $str_to_find = "(" . $matches[1] . ")"; // I tried doing a regex like /\(([0-9+* ]+)\)/ to also capture the ( and ) but it does not seem to work so let"s add them myself
                $string = str_replace($str_to_find, (string)$this->solvePart1($matches[1]), $string);
            }
        }


        // Now we should be able to go left-to-right with no more ( or )
        $items = explode(" ", $string);
        // We need at lest 3 values but we can have more
        while (count($items) > 1) {
            $number1 = (int)array_shift($items);
            $operator = array_shift($items);
            $number2 = (int)array_shift($items);

            // Now calculate, the result will be added to the start of the array so we get something like
            // 1 + 2 + 3 + 4 -> 3 + 3 + 4 -> 6 + 4 -> 10

            $val = match ($operator) {
                "*" => $number1 * $number2,
                "+" => $number1 + $number2,
            };
            array_unshift($items, $val);
        }

        // We now have a single value in our array
        return $items[0];
    }

    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 18, 2);

        return 0;
    }


}
