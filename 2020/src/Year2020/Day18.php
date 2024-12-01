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

    public function solvePart1(string $string): int
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
//            array_unshift($items, eval('return ' . $number1 . $operator . $number2 . ";")); // I can't work on AoC with at least a single eval!
        }

        // We now have a single value in our array
        return $items[0];
    }

    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 18, 2);

        $sum = [];
        foreach ($input as $line) {
            $sum[] = $this->solvePart2($line);
        }

        return array_sum($sum);
    }

    public function solvePart2(string $string): int
    {
        $string = trim($string);
        $this->log($string);

        // Recursive thingy to solve every thing between ( ... )
        while (str_contains($string, "(")) {
            if (preg_match("/\(([0-9+* ]+)\)/", $string, $matches)) {
                $str_to_find = "(" . $matches[1] . ")"; // I tried doing a regex like /\(([0-9+* ]+)\)/ to also capture the ( and ) but it does not seem to work so let"s add them myself
                $string = str_replace($str_to_find, (string)$this->solvePart2($matches[1]), $string);
                $this->log($string);
            }
        }

        // No more ( .. ) here, first handle all additions (again just ugly string replaces)
        while (str_contains($string, "+")) {
            if (preg_match("/(\d+) \+ (\d+)/", $string, $matches)) {
                // I only want to replace the FIRST occurrence! https://stackoverflow.com/a/1252710/2451037
                $pos = strpos($string, $matches[0]);
                if ($pos !== false) {
                    $string = substr_replace($string, (string)($matches[1] + $matches[2]), $pos, strlen($matches[0]));
                }
                $this->log($string);
            }
        }

        // Not only the multiplies are left over!
        while (str_contains($string, "*")) {
            if (preg_match("/(\d+) \* (\d+)/", $string, $matches)) {
                // I only want to replace the FIRST occurrence! https://stackoverflow.com/a/1252710/2451037
                $pos = strpos($string, $matches[0]);
                if ($pos !== false) {
                    $string = substr_replace($string, (string)($matches[1] * $matches[2]), $pos, strlen($matches[0]));
                }
                $this->log($string);
            }
        }

        // We should now have a single item in our string that is numeric (if not I have a bug)
        if (!is_numeric($string)) {
            $this->log("Returned value is not numeric " . $string);
        }
        $this->log("Answer: " . $string);
        $this->log("");

        return (int)$string;
    }
}
