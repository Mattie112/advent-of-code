<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day11 extends Day
{
    public const FLOOR = ".";
    public const SEAT_EMPTY = "L";
    public const SEAT_OCCUPIED = "#";
    public array $seats = [];

    public function part1(): int|string
    {
        $this->seats = [];
        $input = $this->getInputAsArray(2020, 11, 1);
        foreach ($input as $row) {
            $this->seats[] = str_split($row);
        }
        $prev_occupied_seats = 0;
        $tmp_seats = $this->seats;
        while (true) {
            foreach ($this->seats as $row_id => $row) {
                foreach ($row as $seat_id => $seat) {
                    if ($seat === self::FLOOR) {
                        continue;
                    }
                    $checked_seats = [];
                    for ($tmp_row = -1; $tmp_row <= 1; $tmp_row++) {
                        for ($tmp_col = -1; $tmp_col <= 1; $tmp_col++) {
                            if ($tmp_col === 0 && $tmp_row === 0) {
                                continue;
                            }
                            $checked_seats[] = $this->getSeatById($row_id - $tmp_row, $seat_id - $tmp_col);
                        }
                    }

                    switch ($seat) {
                        case self::FLOOR:
                            break;
                        case self::SEAT_EMPTY:
                            // First see if ALL seats are empty
                            if (count(array_filter($checked_seats, static fn($item) => $item === self::SEAT_OCCUPIED)) === 0) {
                                $tmp_seats[$row_id][$seat_id] = self::SEAT_OCCUPIED;
                            }

                            break;
                        case self::SEAT_OCCUPIED:
                            // If >4 seats are taken this seat will be empty
                            if (count(array_filter($checked_seats, static fn($item) => $item === self::SEAT_OCCUPIED)) >= 4) {
                                $tmp_seats[$row_id][$seat_id] = self::SEAT_EMPTY;
                            }
                            break;
                    }
                }
            }
            $this->seats = $tmp_seats;

            if ($this->output?->isVerbose()) {
                $this->debug($this->seats);
            }

            $occupied_seats = 0;
            foreach ($this->seats as $row) {
                $occupied_seats += count(array_filter($row, static fn($item) => $item === self::SEAT_OCCUPIED));
            }
            if ($occupied_seats === $prev_occupied_seats) {
                // No change this round we are stabilized
                return $occupied_seats;
            }
            $prev_occupied_seats = $occupied_seats;
            $this->log($prev_occupied_seats . " - " . $occupied_seats);
        }
    }

    public function debug($seats): void
    {
        foreach ($seats as $row_id => $row) {
            foreach ($row as $seat_id => $seat) {
                echo $seat;
            }
            echo PHP_EOL;
        }
        echo PHP_EOL;
        echo PHP_EOL;
    }

    public function getSeatById(int $row_id, int $seat_id)
    {
        return $this->seats[$row_id][$seat_id] ?? self::FLOOR;
    }

    public function part2(): int|string
    {
        $this->seats = [];
        $input = $this->getInputAsArray(2020, 11, 2);
        foreach ($input as $row) {
            $this->seats[] = str_split($row);
        }
        $prev_occupied_seats = 0;
        $tmp_seats = $this->seats;
        $row_length = count($this->seats);
        $column_length = count($this->seats[0]);

        while (true) {
            foreach ($this->seats as $row_id => $row) {
                foreach ($row as $seat_id => $seat) {
                    if ($seat === self::FLOOR) {
                        continue;
                    }
                    $checked_seats = [];

                    // I have typed-out each for loop so I could so some debugging (note: dont break on a FULL seat but also on an EMPTY one)
                    // I think I might optimise this with a "direction" array like [-1,-1],[-1,0][-1,1] etc

                    // Check our current row left and right
                    for ($tmp_col = $seat_id - 1; $tmp_col >= 0; $tmp_col--) {
                        $seatById = $this->getSeatById($row_id, $tmp_col);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                    }

                    for ($tmp_col = $seat_id + 1; $tmp_col < $column_length; $tmp_col++) {
                        $seatById = $this->getSeatById($row_id, $tmp_col);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                    }

                    // column up
                    for ($tmp_row = $row_id - 1; $tmp_row >= 0; $tmp_row--) {
                        $seatById = $this->getSeatById($tmp_row, $seat_id);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                    }

                    //column down
                    for ($tmp_row = $row_id + 1; $tmp_row < $row_length; $tmp_row++) {
                        $seatById = $this->getSeatById($tmp_row, $seat_id);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                    }

                    // And now we need to check diagonally left/up
                    $tmp_col = $seat_id-1;
                    for ($tmp_row = $row_id - 1; $tmp_row >= 0; $tmp_row--) {
                        $seatById = $this->getSeatById($tmp_row, $tmp_col);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                        $tmp_col--;
                    }

                    // And now we need to check diagonally right/up
                    $tmp_col = $seat_id+1;
                    for ($tmp_row = $row_id - 1; $tmp_row >= 0; $tmp_row--) {
                        $seatById = $this->getSeatById($tmp_row, $tmp_col);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                        $tmp_col++;
                    }


                    // And now we need to check diagonally left/down
                    $tmp_col =$seat_id -1;
                    for ($tmp_row = $row_id + 1; $tmp_row < $row_length; $tmp_row++) {
                        $seatById = $this->getSeatById($tmp_row, $tmp_col);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                        $tmp_col--;
                    }

                    // And now we need to check diagonally right/down
                    $tmp_col = $seat_id+1;
                    for ($tmp_row = $row_id + 1; $tmp_row < $row_length; $tmp_row++) {
                        $seatById = $this->getSeatById($tmp_row, $tmp_col);
                        $checked_seats[] = $seatById;
                        if ($seatById !== self::FLOOR) {
                            break;
                        }
                        $tmp_col++;
                    }

                    switch ($seat) {
                        case self::FLOOR:
                            break;
                        case self::SEAT_EMPTY:
                            // First see if ALL seats are empty
                            if (count(array_filter($checked_seats, static fn($item) => $item === self::SEAT_OCCUPIED)) === 0) {
                                $tmp_seats[$row_id][$seat_id] = self::SEAT_OCCUPIED;
                            }

                            break;
                        case self::SEAT_OCCUPIED:
                            // If >4 seats are taken this seat will be empty
                            if (count(array_filter($checked_seats, static fn($item) => $item === self::SEAT_OCCUPIED)) >= 5) {
                                $tmp_seats[$row_id][$seat_id] = self::SEAT_EMPTY;
                            }
                            break;
                    }
                }
            }
            $this->seats = $tmp_seats;

            if ($this->output?->isVerbose()) {
                $this->debug($this->seats);
            }

            $occupied_seats = 0;
            foreach ($this->seats as $row) {
                $occupied_seats += count(array_filter($row, static fn($item) => $item === self::SEAT_OCCUPIED));
            }
            if ($occupied_seats === $prev_occupied_seats) {
                // No change this round we are stabilized
                return $occupied_seats;
            }
            $prev_occupied_seats = $occupied_seats;
            $this->log($prev_occupied_seats . " - " . $occupied_seats);
        }
    }
}
