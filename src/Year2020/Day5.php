<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day5 extends Day
{
    public const FRONT = "F";
    public const BACK = "B";
    public const LEFT = "L";
    public const RIGHT = "R";

    public function part1(): int|string
    {
        $input = $this->getInput(2020, 5, 1);
        $input = explode("\n", $input);

        $max = 0;
        foreach ($input as $pass) {
            [$row, $col] = $this->findSeat($pass);
            $max = max($max, (($row * 8) + $col));

        }
        return $max;
    }

    public function part2(): int|string
    {
        $input = $this->getInput(2020, 5, 1);
        $input = explode("\n", $input);

        $seat_ids = [];

        foreach ($input as $pass) {
            [$row, $col] = $this->findSeat($pass);
            if ($row === 0 || $row === 127) {
                continue;
            }
            $seat_id = ($row * 8) + $col;
            $seat_ids[$seat_id] = [$row, $col];
        }

        // Now look for empty seats by ID
        $highest_seat_id = max(array_keys($seat_ids));
        $my_seat = PHP_INT_MAX;
        for ($i = 0; $i < $highest_seat_id; $i++) {
            // We can ignore the first row (<= 16 for seat ID; 1 * 8 + 8), we will also ignore the last row by using the lowest seat ID as our answer
            if ($i <= 16) {
                continue;
            }
            if (!isset($seat_ids[$i])) {
                $this->log("Found seat with no boarding pass with ID #" . $i);
                $my_seat = min($my_seat, $i);
            }
        }

        return $my_seat;
    }

    protected function findSeat($pass): array
    {
        $seat_low = 0;
        $seat_high = 127;
        $modifier = 64; // Start with 128 options, but first modification is 64
        $this->log($pass);
        $this->log("We are now at [" . $seat_low . " - " . $seat_high . "]");
        $letters = str_split($pass);

        $checks = 1;
        foreach ($letters as $letter) {
            if ($checks > 7) {
                if ($seat_low !== $seat_high) {
                    // This is strange we should have found a seat by now
                    $this->log("<error>Should have found a seat after 7 iterations</error>");
                }
                break;
            }

            switch ($letter) {
                case self::FRONT;
                    $seat_high -= $modifier;
                    break;
                case self::BACK;
                    $seat_low += $modifier;
                    break;
            }
            $this->log("We are now at [" . $seat_low . " - " . $seat_high . "]");
            $modifier /= 2;
            $checks++;
        }

        // Now we can check the last 3 characters (columns)
        $modifier = 4; // Start with 8 options, but first modification is 4
        $row = $seat_high;
        $seat_low = 0;
        $seat_high = 7;
        $letters = [$letters[7], $letters[8], $letters[9]];
        $this->log("We are now at #" . $row . " [" . $seat_low . " - " . $seat_high . "]");
        foreach ($letters as $letter) {
            switch ($letter) {
                case self::LEFT;
                    $seat_high -= $modifier;
                    break;
                case self::RIGHT;
                    $seat_low += $modifier;
                    break;
            }
            $this->log("We are now at #" . $row . " [" . $seat_low . " - " . $seat_high . "]");
            $modifier /= 2;
        }
        $col = $seat_high;
        return [$row, $col];
    }
}
