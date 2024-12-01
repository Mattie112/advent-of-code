<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day2 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInput(2020, 2, 1);
        $input = explode("\n", $input);
        $valid = 0;
        foreach ($input as $item) {
            if (trim($item) === "") {
                continue;
            }
            preg_match("@(.*)-(.*) ([a-z]): (.*)@", $item, $matches);
            [, $minamount, $maxamount, $letter, $password] = $matches;

            $passwordarr = str_split($password);
            $a = array_filter($passwordarr, static function ($p) use ($letter) {
                return $p === $letter;
            });
            $count = count($a);
            if ($count >= $minamount && $count <= $maxamount) {
                $valid++;
            }
        }
        return $valid;
    }

    public function part2(): int|string
    {
        $input = $this->getInput(2020, 2, 1);
        $input = explode("\n", $input);
        $valid = 0;
        foreach ($input as $item) {
            if (trim($item) === "") {
                continue;
            }
            preg_match("@(.*)-(.*) ([a-z]): (.*)@", $item, $matches);
            [, $minamount, $maxamount, $letter, $password] = $matches;

            $passwordarr = str_split($password);

            if ($passwordarr[$minamount - 1] === $letter xor $passwordarr[$maxamount - 1] === $letter) {
                $valid++;
            }
        }
        return $valid;
    }
}
