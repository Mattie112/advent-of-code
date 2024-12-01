<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day24 extends Day
{
    public const WHITE = 0;
    public const BLACK = 1;
    public int $n = 0;
    public int $e = 0;
    public int $s = 0;
    public int $w = 0;

    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 24, 2);
        $grid = $this->getInitialState($input);
        return $this->calcBlacks($grid);
    }

    public function calcBlacks(array $grid): int
    {
        $count = 0;
        foreach ($grid as $x) {
            foreach ($x as $y) {
                if ($y === self::BLACK) {
                    $count++;
                }
            }
        }

        return $count;
    }

    public function getInitialState(array $input): array
    {
        $grid = [];
        foreach ($input as $line) {
            $y = 0;
            $x = 0;
            while ($line !== '') {
                $next = 0;
                $elem = $line[0];
                switch ($elem) {
                    case "e":
                        // East is always a single character
                        $x += 2;
                        $next = 1;
                        break;
                    case "w":
                        // West is always a single character
                        $x -= 2;
                        $next = 1;
                        break;
                    case "n":
                        // This could be northeast or northwest, here we determine the Y, in the next step the X
                        $y++;
                        $next = 2;
                        break;
                    case "s":
                        // This could be southeast or southwest, here we determine the Y, in the next step the X
                        $y--;
                        $next = 2;
                        break;
                }
                if ($next === 2) {
                    switch ($line[1]) {
                        case "e":
                            // This could be northeast or southeast, now the X
                            $x++;
                            break;
                        case "w":
                            // This could be northwest or southwest, now the X
                            $x--;
                            break;
                    }
                }
                $line = substr($line, $next);
            }

            if (isset($grid[$x][$y]) && $grid[$x][$y] === self::BLACK) {
                // We start with everything "black"
                $grid[$x][$y] = self::WHITE;
            } else {
                $grid[$x][$y] = self::BLACK;
            }

            // Hackyhacky so I know what bound to use for part 2
            $this->e = max($this->e, $x);
            $this->w = min($this->w, $x);
            $this->n = max($this->n, $y);
            $this->s = min($this->s, $y);
        }
        return $grid;
    }


    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 24, 2);
        $grid = $this->getInitialState($input);

        for ($i = 1; $i <= 100; $i++) {
            $copy = $grid;

            // Grow the grid
            --$this->s;
            ++$this->n;
            $this->w -= 2;
            $this->e += 2;

            for ($y = $this->s; $y <= $this->n; $y++) {
                for ($x = $this->w; $x <= $this->e; $x++) {
                    $black_count = 0;
                    // Just like part 1 this is the entire EAST/WEST and NorthEast NorthWest SouthEast SouthWest magic
                    $black_count += (isset($grid[$x - 2][$y + 0]) && $grid[$x - 2][$y + 0] === self::BLACK) ? 1 : 0;
                    $black_count += (isset($grid[$x + 2][$y + 0]) && $grid[$x + 2][$y + 0] === self::BLACK) ? 1 : 0;
                    $black_count += (isset($grid[$x - 1][$y - 1]) && $grid[$x - 1][$y - 1] === self::BLACK) ? 1 : 0;
                    $black_count += (isset($grid[$x + 1][$y - 1]) && $grid[$x + 1][$y - 1] === self::BLACK) ? 1 : 0;
                    $black_count += (isset($grid[$x - 1][$y + 1]) && $grid[$x - 1][$y + 1] === self::BLACK) ? 1 : 0;
                    $black_count += (isset($grid[$x + 1][$y + 1]) && $grid[$x + 1][$y + 1] === self::BLACK) ? 1 : 0;
                    if (isset($grid[$x][$y]) && $grid[$x][$y] === self::BLACK) {
                        if ($black_count === 0 || $black_count > 2) {
                            $copy[$x][$y] = self::WHITE;
                        }
                    } else if ($black_count === 2) {
                        $copy[$x][$y] = self::BLACK;
                    }
                }
            }
            $grid = $copy;
            if ($this->output?->isVerbose()) {
                $this->log($this->calcBlacks($grid));
            }
        }

        return $this->calcBlacks($grid);
    }
}
