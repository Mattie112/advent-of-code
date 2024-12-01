<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day8 extends Day
{
    protected int $acc = 0;

    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 8, 1);

        $this->executeProgram($input);
        return $this->acc;
    }

    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 8, 2);
        // Go through the code again and for each line check for a nop or jmp, change it and execute it to see if it works
        foreach ($input as $i => $line) {
            preg_match("@(.{3})@", $input[$i], $matches);
            $operation = $matches[1];
            $copy = $input;
            // Let's try out a new PHP8 thingy
            $copy[$i] = match ($operation) {
                "acc" => $line,
                "nop" => str_replace("nop", "jmp", $line),
                "jmp" => str_replace("jmp", "nop", $line),
            };

            // Execute program and on a clean exit return the accumulator (= answer for part 2)
            if ($this->executeProgram($copy)) {
                return $this->acc;
            }
        }
        return 0;
    }

    public function executeProgram(array $code): bool
    {
        $this->acc = 0;
        $already_visited = [];
        $linecount = count($code);
        for ($i = 0; $i < $linecount;) {
            preg_match("@(.{3}) ([+-]\d+)@", $code[$i], $matches);
            $operation = $matches[1];
            $argument = (int)$matches[2];

            switch ($operation) {
                case "acc":
                    $this->acc += $argument;
                    $i++;
                    break;
                case "jmp":
                    $i += $argument;
                    break;
                case "nop":
                    $i++;
                    break;
            }

            // And right before we go to execute the next one (our $i is already changed here we check if we have execute that line before
            if (isset($already_visited[$i])) {
                // If that is the case return false (=endless loop)
                return false;
            }
            $already_visited[$i] = true;
        }
        return true;
    }
}
