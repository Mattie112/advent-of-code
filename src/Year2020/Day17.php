<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day17 extends Day
{
    public int $startW = 0;
    public int $startZ = 0;
    public int $startY = 0;
    public int $startX = 0;

    // My grid is (W) Z Y X

    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 17, 1);
        $grid = [];
        foreach ($input as $y => $line) {
            $line = str_split($line);
            foreach ($line as $x => $char) {
                if ($char === '#') {
                    $grid[0][$y][$x] = true;
                } else {
                    $grid[0][$y][$x] = false;
                }
            }
        }

        // So we know how much to grow
        $this->startX = (int)(0 - floor(strlen(trim($input[0]))));
        $this->startY = (int)(0 - floor(strlen(trim($input[0]))));

        $this->log(sprintf("Before cycle %d I have %d active", 0, $this->countActive($grid)));

        for ($cycle = 1; $cycle <= 6; $cycle++) {
            $this->startX--;
            $this->startY--;
            $this->startZ--;
            for ($z = $this->startZ; $z <= -1 * $this->startZ; $z++) {
                for ($y = $this->startY; $y <= -1 * $this->startY; $y++) {
                    for ($x = $this->startX; $x <= -1 * $this->startX; $x++) {
                        if (!isset($grid[$z][$y][$x])) {
                            $grid[$z][$y][$x] = false;
                        }
                    }
                }
            }

            // Make the debugging easier
            $this->recursive_ksort($grid);

            $copy = $grid;

            foreach ($copy as $z => $zarr) {
                foreach ($zarr as $y => $yarr) {
                    foreach ($yarr as $x => $cube) {
                        // Now chech the neighbours
                        $count = $this->checkNeighbours($z, $y, $x, $copy);
                        if ($cube === true && ($count === 3 || $count === 2)) {
                            // Cube remains active (added for debugging)
                            $grid[$z][$y][$x] = true;
                        } else {
                            $grid[$z][$y][$x] = false;
                        }
                        if ($cube === false && $count === 3) {
                            // Cube becomes active
                            $grid[$z][$y][$x] = true;
                        }
                    }
                }
            }

            $this->log(sprintf("After cycle %d I have %d active", $cycle, $this->countActive($grid)));
        }

        return $this->countActive($grid);
    }


    public function recursive_ksort(&$array): bool
    {
        foreach ($array as $k => &$v) {
            if (is_array($v)) {
                $this->recursive_ksort($v);
            }
        }
        return ksort($array);
    }

    public function countActive(array $grid): int
    {
        $count_active = 0;
        foreach ($grid as $z => $zarr) {
            foreach ($zarr as $y => $yarr) {
                foreach ($yarr as $x => $cube) {
                    if ($cube === true) {
                        $count_active++;
                    }
                }
            }
        }

        return $count_active;
    }

    public function checkNeighbours(int $z, int $y, int $x, array $grid): int
    {
        $count = 0;
        for ($iz = $z - 1; $iz <= $z + 1; $iz++) {
            for ($iy = $y - 1; $iy <= $y + 1; $iy++) {
                for ($ix = $x - 1; $ix <= $x + 1; $ix++) {
                    if ($ix === $x && $iy === $y && $iz === $z) {
                        continue;
                    }
                    if (isset($grid[$iz][$iy][$ix]) && $grid[$iz][$iy][$ix] === true) {
                        $count++;
                    }
                }
            }
        }
        return $count;
    }

    // Copy from part 1 with a $w added
    public function countActive4D(array $grid): int
    {
        $count_active = 0;
        foreach ($grid as $w => $warr) {
            foreach ($warr as $z => $zarr) {
                foreach ($zarr as $y => $yarr) {
                    foreach ($yarr as $x => $cube) {
                        if ($cube === true) {
                            $count_active++;
                        }
                    }
                }
            }
        }

        return $count_active;
    }

    // Copy from part 1 with a $w added
    public function checkNeighbours4D(int $w, int $z, int $y, int $x, array $grid): int
    {
        $count = 0;
        for ($iw = $w - 1; $iw <= $w + 1; $iw++) {
            for ($iz = $z - 1; $iz <= $z + 1; $iz++) {
                for ($iy = $y - 1; $iy <= $y + 1; $iy++) {
                    for ($ix = $x - 1; $ix <= $x + 1; $ix++) {
                        if ($iw === $w && $ix === $x && $iy === $y && $iz === $z) {
                            continue;
                        }
                        if (isset($grid[$iw][$iz][$iy][$ix]) && $grid[$iw][$iz][$iy][$ix] === true) {
                            $count++;
                        }
                    }
                }
            }
        }
        return $count;
    }

    // Copy from part 1 with a $w added
    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 17, 2);
        $grid = [];
        foreach ($input as $y => $line) {
            $line = str_split($line);
            foreach ($line as $x => $char) {
                if ($char === '#') {
                    $grid[0][0][$y][$x] = true;
                } else {
                    $grid[0][0][$y][$x] = false;
                }
            }
        }

        // So we know how much to grow
        $this->startX = (int)(0 - floor(strlen(trim($input[0]))));
        $this->startY = (int)(0 - floor(strlen(trim($input[0]))));

        $this->log(sprintf("Before cycle %d I have %d active", 0, $this->countActive4D($grid)));

        for ($cycle = 1; $cycle <= 6; $cycle++) {
            $this->startX--;
            $this->startY--;
            $this->startZ--;
            $this->startW--;
            for ($w = $this->startW; $w <= -1 * $this->startW; $w++) {
                for ($z = $this->startZ; $z <= -1 * $this->startZ; $z++) {
                    for ($y = $this->startY; $y <= -1 * $this->startY; $y++) {
                        for ($x = $this->startX; $x <= -1 * $this->startX; $x++) {
                            if (!isset($grid[$w][$z][$y][$x])) {
                                $grid[$w][$z][$y][$x] = false;
                            }
                        }
                    }
                }
            }

            // Make the debugging easier
            $this->recursive_ksort($grid);

            $copy = $grid;

            foreach ($copy as $w => $warr) {
                foreach ($warr as $z => $zarr) {
                    foreach ($zarr as $y => $yarr) {
                        foreach ($yarr as $x => $cube) {
                            // Now chech the neighbours
                            $count = $this->checkNeighbours4D($w, $z, $y, $x, $copy);
                            if ($cube === true && ($count === 3 || $count === 2)) {
                                // Cube remains active (added for debugging)
                                $grid[$w][$z][$y][$x] = true;
                            } else {
                                $grid[$w][$z][$y][$x] = false;
                            }
                            if ($cube === false && $count === 3) {
                                // Cube becomes active
                                $grid[$w][$z][$y][$x] = true;
                            }
                        }
                    }
                }
            }

            $this->log(sprintf("After cycle %d I have %d active", $cycle, $this->countActive4D($grid)));
        }

        return $this->countActive4D($grid);
    }
}
