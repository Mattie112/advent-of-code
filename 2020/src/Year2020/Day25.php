<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day25 extends Day
{
    public const MAGIC_NUMBER = 20201227;
    public const SUBJECT_NUMBER = 7;

    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 25, 1);
        $keys = array_map("intval", $input);

        $loops[] = $this->findLoopSize($keys[0]);
        $loops[] = $this->findLoopSize($keys[1]);

        return $this->getEncryptionKey($keys[0], $loops[1]); // or keys[1] and loops[0]
    }

    public function findLoopSize(int $key): int
    {
        $value = 1;
        $loop = 0;
        while ($value !== $key) {
            $loop++;
            $value = ($value * self::SUBJECT_NUMBER) % self::MAGIC_NUMBER;
        }
        return $loop;
    }

    public function getEncryptionKey(int $card_number, int $loops): int
    {
        $value = 1;
        for ($i = 0; $i < $loops; $i++) {
            $value = ($value * $card_number) % self::MAGIC_NUMBER;
        }
        return $value;
    }

    public function part2(): int|string
    {
        $input = $this->getInput(2020, 25, 2);

        return 0;
    }


}
