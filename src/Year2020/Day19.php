<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day19 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 19, 1, PHP_EOL . PHP_EOL);
        $raw_rules = explode(PHP_EOL, $input[0]);
        $data_lines = explode(PHP_EOL, $input[1]);

        $rules = [];
        foreach ($raw_rules as $i => $rule) {
            $a = 1;
            if (preg_match("/(\d+): \"([ab])/", $rule, $matches)) {
                $rules[(int)$matches[1]] = $matches[2];
            } else if (preg_match("/(\d+): (\d+) (\d+) \| (\d+) (\d+)$/", $rule, $matches)) {
                // Match stuff with the OR pipe (1 2 | 3 4)
                $rules[(int)$matches[1]][] = [(int)$matches[2], (int)$matches[3]];
                $rules[(int)$matches[1]][] = [(int)$matches[4], (int)$matches[5]];
            } else if (preg_match("/(\d+): (\d+) \| (\d+)$/", $rule, $matches)) {
                // Match stuff with the OR pipe (1 | 2)
//                $rules[(int)$matches[1]][] = [(int)$matches[2], (int)$matches[3]];
                $rules[(int)$matches[1]][] = [(int)$matches[2]];
                $rules[(int)$matches[1]][] = [(int)$matches[3]];
            } else if (preg_match("/(\d+): (\d+) (\d+) ?(\d)?$/", $rule, $matches)) {
                // Only the tests uses 3 numbers so we hack this in
                $tmp[] = (int)$matches[2];
                $tmp[] = (int)$matches[3];
                if (isset($matches[4])) {
                    $tmp[] = (int)$matches[4];
                    $this->log("only test " . $i);
                }
                $rules[(int)$matches[1]][] = $tmp;
            } elseif (preg_match("/(\d+): (\d+)$/", $rule, $matches)) {
                // Match single digit
                $rules[(int)$matches[1]] = (int)$matches[2];
            } else {
                $this->log("Could not regex match " . $rule);
            }
            unset($tmp);
        }
        ksort($rules); // Easy debugging

        $regex = trim($this->findRuleValue($rules, 0));
        $this->log($regex);

        $valid = 0;
        foreach ($data_lines as $data) {
            if (preg_match("/^" . $regex . "$/", $data)) {
                $valid++;
            }
        }

        return $valid;
    }

    public function findRuleValue(array $rules, int $id): string
    {
        $rule = "";
        $rule_to_check = $rules[$id];
        if ($rule_to_check === "a" || $rule_to_check === "b") {
            return $rule_to_check;
        } else if (is_array($rule_to_check)) {
            $a = 1;
            if (count($rule_to_check) > 1) {
                // we have an OR thingy
                $tmp = "(?:";
                foreach ($rule_to_check[0] as $tt) {
                    $tmp .= $this->findRuleValue($rules, $tt);
                }
                $tmp .= "|";
                foreach ($rule_to_check[1] as $tt) {
                    $tmp .= $this->findRuleValue($rules, $tt);
                }
                $tmp .= ")";
                $rule .= $tmp;
            } else if (count($rule_to_check) === 1) {
                foreach ($rule_to_check[0] as $tt) {
                    $rule .= $this->findRuleValue($rules, $tt);
                }
            } else {
                $this->log("to many");
            }
        } elseif (is_numeric($rule_to_check)) {
            $rule .= $this->findRuleValue($rules, $rule_to_check);
        } else {
            $this->log("wtf is this?");
        }
        $this->log($rule);
        return $rule;
    }

    public function part2(): int|string
    {
        // todo run this and check the console, you'll see repetition, todo decide on how to find/fix this. Perhaps simply store all found rules or something?

        ini_set("memory_limit", "5G");
        $input = $this->getInputAsArray(2020, 19, 2, PHP_EOL . PHP_EOL);
        $raw_rules = explode(PHP_EOL, $input[0]);
        $data_lines = explode(PHP_EOL, $input[1]);

        $rules = [];
        foreach ($raw_rules as $i => $rule) {
            $a = 1;
            if (preg_match("/(\d+): \"([ab])/", $rule, $matches)) {
                $rules[(int)$matches[1]] = $matches[2];
            } else if (preg_match("/(\d+): (\d+) (\d+) \| (\d+) (\d+) (\d+)$/", $rule, $matches)) {
                // Match stuff with the OR pipe (1 2 | 3 4)
                $rules[(int)$matches[1]][] = [(int)$matches[2], (int)$matches[3]];
                $rules[(int)$matches[1]][] = [(int)$matches[4], (int)$matches[5], (int)$matches[6]];
            } else if (preg_match("/(\d+): (\d+) \| (\d+) (\d+)$/", $rule, $matches)) {
                // Match stuff with the OR pipe (1 | 3 4)
                $rules[(int)$matches[1]][] = [(int)$matches[2]];
                $rules[(int)$matches[1]][] = [(int)$matches[3], (int)$matches[4]];
            } else if (preg_match("/(\d+): (\d+) (\d+) \| (\d+) (\d+)$/", $rule, $matches)) {
                // Match stuff with the OR pipe (1 2 | 3 4)
                $rules[(int)$matches[1]][] = [(int)$matches[2], (int)$matches[3]];
                $rules[(int)$matches[1]][] = [(int)$matches[4], (int)$matches[5]];
            } else if (preg_match("/(\d+): (\d+) \| (\d+)$/", $rule, $matches)) {
                // Match stuff with the OR pipe (1 | 2)
//                $rules[(int)$matches[1]][] = [(int)$matches[2], (int)$matches[3]];
                $rules[(int)$matches[1]][] = [(int)$matches[2]];
                $rules[(int)$matches[1]][] = [(int)$matches[3]];
            } else if (preg_match("/(\d+): (\d+) (\d+) ?(\d)?$/", $rule, $matches)) {
                // Only the tests uses 3 numbers so we hack this in
                $tmp[] = (int)$matches[2];
                $tmp[] = (int)$matches[3];
                if (isset($matches[4])) {
                    $tmp[] = (int)$matches[4];
                    $this->log("only test " . $i);
                }
                $rules[(int)$matches[1]][] = $tmp;
            } elseif (preg_match("/(\d+): (\d+)$/", $rule, $matches)) {
                // Match single digit
                $rules[(int)$matches[1]] = (int)$matches[2];
            } else {
                $this->log("Could not regex match " . $rule);
            }
            unset($tmp);
        }
        ksort($rules); // Easy debugging

        $regex = trim($this->findRuleValue($rules, 0));
        $this->log($regex);

        $valid = 0;
        foreach ($data_lines as $data) {
            if (preg_match("/^" . $regex . "$/", $data)) {
                $valid++;
            }
        }

        return $valid;
    }
}
