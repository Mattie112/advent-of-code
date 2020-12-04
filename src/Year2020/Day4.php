<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day4 extends Day
{

    public function part1(): int|string
    {
        $input = $this->getUntrimmedInput(2020, 4, 1, false);
        $input = explode("\n", $input);
        $input[] = "";

        $passports = [];
        $temp = "";
        foreach ($input as $line) {
            // Go ahead and regex
            if ($line === "") {
                $passports[] = $temp;
                $temp = "";
            }
            $temp .= " " . trim($line, "\n");
        }

        $valid = 0;
        foreach ($passports as $passport) {
            $byr = null;
            $iyr = null;
            $eyr = null;
            $hgt = null;
            $hcl = null;
            $ecl = null;
            $pid = null;
            $cid = null;
            preg_match("@byr:(.+)@", $passport, $a);
            if ($a) {
                $byr = $a[0];
            }

            preg_match("@iyr:(.+)@", $passport, $a);
            if ($a) {
                $iyr = $a[0];
            }

            preg_match("@eyr:(.+)@", $passport, $a);
            if ($a) {
                $eyr = $a[0];
            }

            preg_match("@hgt:(.+)@", $passport, $a);
            if ($a) {
                $hgt = $a[0];
            }

            preg_match("@hcl:(.+)@", $passport, $a);
            if ($a) {
                $hcl = $a[0];
            }

            preg_match("@ecl:(.+)@", $passport, $a);
            if ($a) {
                $ecl = $a[0];
            }

            preg_match("@pid:(.+)@", $passport, $a);
            if ($a) {
                $pid = $a[0];
            }

            preg_match("@cid:(.+)@", $passport, $a);
            if ($a) {
                $cid = $a[0];
            }

            if ($byr !== null && $iyr !== null && $eyr !== null && $hgt !== null && $hcl !== null && $ecl !== null && $pid !== null) {
                $valid++;
            }
        }
        return $valid;
    }

    public function part2(): int|string
    {
        $input = $this->getUntrimmedInput(2020, 4, 2, false);
        $input = explode("\n", $input);
        $input[] = "";

        $passports = [];
        $temp = "";
        foreach ($input as $line) {
            // Go ahead and regex
            if ($line === "") {
                $passports[] = $temp;
                $temp = "";
            }
            $temp .= " " . str_replace("\n", " ", $line) . " ";
        }


        $valid = 0;
        foreach ($passports as $passport) {
            $byr = null;
            $iyr = null;
            $eyr = null;
            $hgt = null;
            $hcl = null;
            $ecl = null;
            $pid = null;
            $cid = null;
            preg_match("@byr:(\d{4})\s@", $passport, $a);
            if ($a) {
                if ($a[1] >= 1920 && $a[1] <= 2002) {
                    $byr = $a[1];
                }
            }

            preg_match("@iyr:(\d{4})\s@", $passport, $a);
            if ($a) {
                if ($a[1] >= 2010 && $a[1] <= 2020) {
                    $iyr = $a[1];
                }
            }

            preg_match("@eyr:(\d{4})\s@", $passport, $a);
            if ($a) {
                if ($a[1] >= 2020 && $a[1] <= 2030) {
                    $eyr = $a[1];
                }
            }

            preg_match("@hgt:((\d+)(in|cm))\s@", $passport, $a);
            if ($a) {
                if (str_contains($a[1], "cm")) {
                    if ($a[2] >= 150 && $a[2] <= 193) {
                        $hgt = $a[2];
                    }
                } else {
                    if ($a[2] >= 59 && $a[2] <= 76) {
                        $hgt = $a[2];
                    }
                }
            }

            preg_match("@hcl:#([a-f0-9]{6})\s@", $passport, $a);
            if ($a) {
                $hcl = $a[1];
            }

            preg_match("@ecl:(amb|blu|brn|gry|grn|hzl|oth)\s@", $passport, $a);
            if ($a) {
                $ecl = $a[1];
            }

            preg_match("@pid:(\d{9})\s@", $passport, $a);
            if ($a) {
                $pid = $a[1];
            }

            preg_match("@cid:(\d+)\s@", $passport, $a);
            if ($a) {
                $cid = $a[1];
            }
            if ($byr !== null && $iyr !== null && $eyr !== null && $hgt !== null && $hcl !== null && $ecl !== null && $pid !== null) {
                $valid++;
            }
        }
        return $valid;
    }
}
