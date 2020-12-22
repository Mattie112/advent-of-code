<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day21 extends Day
{

    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 21, 1);

        $ingredients_with_possible_allergens = [];
        foreach ($input as $line) {
            $ingredients_tmp = substr($line, 0, strpos($line, "("));
            $allergens_tmp = substr($line, strpos($line, "contains ") + 9);

            $tmp = [];
            if (preg_match_all("/(?:([a-z]+) ?)/", $ingredients_tmp, $matches)) {
                foreach ($matches[1] as $m) {
                    if(!isset($ingredients_with_possible_allergens[$m])) {
                        $ingredients_with_possible_allergens[$m] = [];
                    }
                    $tmp[$m] = [];
                }
            }
            if (preg_match_all("/(?:([a-z]+) ?)/", $allergens_tmp, $matches)) {
                foreach ($tmp as $t => $_) {
                    foreach ($matches[1] as $m) {
                        $ingredients_with_possible_allergens[$t][] = $m;
                    }
                }
            }
        }


        $andersom = [];
        foreach ($input as $line) {
            $ingredients_tmp = substr($line, 0, strpos($line, "("));
            $allergens_tmp = substr($line, strpos($line, "contains ") + 9);

            $tmp = [];
            if (preg_match_all("/(?:([a-z]+) ?)/", $allergens_tmp, $matches)) {
                foreach ($matches[1] as $m) {
                    if(!isset($andersom[$m])) {
                        $andersom[$m] = [];
                    }
                    $tmp[$m] = [];
                }

            }
            if (preg_match_all("/(?:([a-z]+) ?)/", $ingredients_tmp, $matches)) {
                foreach ($tmp as $t => $_) {
                    foreach ($matches[1] as $m) {
                        $andersom[$t][] = $m;
                    }
                }
            }

        }

        return 0;
    }

    public function part2(): int|string
    {
        return 0;

    }

}
