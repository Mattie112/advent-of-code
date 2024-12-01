<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day7 extends Day
{
    protected array $all_bags = [];

    /*
     * These solutions are a bit messy, I could not get the correct single regex to parse it in one go and as I spend to much time on this I'm gonna leave it the way it is :)
     */

    public function part1(): int|string
    {
        $input = $this->getInput(2020, 7, 1);
        $input = explode("\n", $input);

        $this->all_bags = [];
        foreach ($input as $line) {
            preg_match_all("@(([\w]+ ){2})bag@", $line, $matches);

            if (str_contains($line, "contain no other bags")) {
                $this->all_bags[$matches[1][0]] = [];
                continue;
            }

            // Now any remaining matches contain children
            $children = array_slice($matches[1], 1);
            $this->all_bags[$matches[1][0]] = $children;
        }

        $counter = 0;
        foreach ($this->all_bags as $parent_bag => $children) {
            $amount = $this->findGoldenBag($children);
            if ($amount >= 1) {
                $counter++;
            }
        }

        return $counter;
    }

    public function findGoldenBag(array $children): int
    {
        $counter = 0;
        foreach ($children as $child_bag) {
            if ($child_bag === "shiny gold ") {
                return 1;
            }
            $counter += $this->findGoldenBag($this->all_bags[$child_bag]);
        }
        return $counter;
    }

    public function part2(): int|string
    {
        $input = $this->getInput(2020, 7, 1);
        $input = explode("\n", $input);

        foreach ($input as $line) {
            preg_match_all("@(\d+)? ?(([\w]+ ){2})bag@", $line, $matches);

            if (str_contains($line, "contain no other bags")) {
                $this->all_bags[$matches[2][0]] = [];
                continue;
            }

            // Now any remaining matches contain children
            $child_keys = array_slice($matches[2], 1);
            $child_values = array_slice($matches[1], 1);
            $child_values = array_map(static function ($elem) {
                return (int)$elem;
            }, $child_values);
            $children = array_combine($child_keys, $child_values);
            $this->all_bags[$matches[2][0]] = $children;
        }

        return $this->countBagsRecursive($this->all_bags['shiny gold ']);
    }

    public function countBagsRecursive(array $children)
    {
        $counter = 0;
        foreach ($children as $child_bag => $value) {
            $sum = $this->countBagsRecursive($this->all_bags[$child_bag]);
            $counter += $value + ($value * $sum);
        }

        return $counter;
    }

}
