<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day8 extends Day
{


    public function part1(): int|string
    {
        $input = $this->getInput(2020, 8, 1);
        $input = explode("\n", $input);
        $already_visited = [];
        $acc = 0;

        $linecount = count($input);
        $i = 0;
        while ($i < $linecount) {
            preg_match("@(.{3}) ([+-]\d+)@", $input[$i], $matches);
            $operation = $matches[1];
            $argument = (int)$matches[2];

            switch ($operation) {
                case "acc":
                    $acc += $argument;
                    $i++;
                    break;
                case "jmp":
                    $i += $argument;
                    break;
                case "nop":
                    $i++;
                    break;
            }

            // And right before we go to execute the next one we check if we had that one already
            if (isset($already_visited[$i])) {
                break;
            }

            $already_visited[$i] = true;
        }

        return $acc;
    }


    public function part2(): int|string
    {
        $input = $this->getInput(2020, 8, 1);
        $input = explode("\n", $input);
        $already_visited = [];
        $acc = 0;

        $linecount = count($input);
        $repair_pointer = -1;
        $repair = false;
        $original = "";
        while (true) {
            // Loop forever until we have it repaired
            if ($repair) {
                foreach ($input as $id => $line) {
                    if ($id <= $repair_pointer) {
                        continue;
                    }

                    preg_match("@(.{3})@", $line, $matches);
                    $operation = $matches[1];
                    switch ($operation) {
                        case "jmp":
                            $original = $line;
                            $input[$id] = str_replace("jmp", "nop", $line);
                            $repair_pointer = $id;
                            break 2;
                        case "nop":
                            $original = $line;
                            $input[$id] = str_replace("nop", "jmp", $line);
                            $repair_pointer = $id;
                            break 2;
                    }
                }
            }
            $repair = false;

            $i = 0;
            while ($i < $linecount) {
                preg_match("@(.{3}) ([+-]\d+)@", $input[$i], $matches);
                $operation = $matches[1];
                $argument = (int)$matches[2];

                switch ($operation) {
                    case "acc":
                        $acc += $argument;
                        $i++;
                        break;
                    case "jmp":
                        $i += $argument;
                        break;
                    case "nop":
                        $i++;
                        break;
                }

                // And right before we go to execute the next one we check if we had that one already
                if (isset($already_visited[$i])) {
                    if(!empty($original)){
                        $input[$repair_pointer] = $original;
                    }
                    $already_visited = [];
                    $acc = 0;
                    $repair = true;
                    break;
                }

                if ($i === $linecount) {
                    return $acc;
                }

                $already_visited[$i] = true;
            }

        }

        return $acc;
    }


}
