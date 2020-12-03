<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day3 extends Day
{
    public const TREE = "#";

    public function part1(): int|string
    {
        $input = $this->parseInput($this->getInput(2020, 3, 1, false));
        return $this->slipperySlope($input, 3, 1);
    }

    public function part2(): int|string
    {
        $input = $this->parseInput($this->getInput(2020, 3, 2, false));

        $a[] = $this->slipperySlope($input, 1, 1);
        $a[] = $this->slipperySlope($input, 3, 1);
        $a[] = $this->slipperySlope($input, 5, 1);
        $a[] = $this->slipperySlope($input, 7, 1);
        $a[] = $this->slipperySlope($input, 1, 2);

        return array_product($a);
    }

    protected function parseInput(string $input): array
    {
        $parsed_input = explode("\n", $input);
        $parsed_input = array_map(static function ($row) {
            return str_split($row);
        }, $parsed_input);
        return $parsed_input;
    }

    public function slipperySlope($input, int $right, int $down): int
    {
        $lines = count($input);
        $cols = count($input[0]);
        $x = 0;
        $y = 0;
        $trees = 0;
        for ($i = 0; $i < $lines; $i++) {
            if (isset($input[$y][$x % $cols]) && $input[$y][$x % $cols] === self::TREE) {
                $trees++;
            }
            $x += $right;
            $y += $down;
        }

        return $trees;
    }
}
