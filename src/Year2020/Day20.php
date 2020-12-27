<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day20 extends Day
{
    public const LEFT = 1;
    public const RIGHT = 2;
    public const UP = 3;
    public const DOWN = 4;

    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 20, 1, PHP_EOL . PHP_EOL);

        /** @var Tile[] $tiles */
        $tiles = [];

        foreach ($input as $tile) {
            $tile = explode(PHP_EOL, $tile);
            $tile_id = 0;
            if (preg_match("/Tile (\d+):/", $tile[0], $matches)) {
                $tile_id = (int)$matches[1];
            }
            unset($tile[0]);
            // Now the tile itself
            $tile = new Tile($tile_id, array_values(array_map("str_split", $tile)));
            $tiles[$tile_id] = $tile;
        }

        // Start with the first tile as position 0,0
        /** @var Tile[Tile] $grid */
        $grid = [[]];
        $grid[0][0] = array_pop($tiles);

        while (count($tiles) > 0) {
            // Now go through our unmatched tiles and try to match it in the grid
            foreach ($tiles as $unmatched_id => $unmatched_tile) {
                foreach ($grid as $y => $yarr) {
                    foreach ($yarr as $x => $grid_tile) {
                        /** @var $grid_tile Tile */
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);
                            $this->log(sprintf("Just the way it is %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }

                        // - MAJOR COPY PASTE ALERT PERHAPS I WILL MAKE THIS A BIT NICER BUT HEY.... IT's TIME FOR LUNCH!
                        $unmatched_tile->rotateRight();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right flip cols %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipRows();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right flip cols and rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right flip rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        // -

                        // -
                        $unmatched_tile->rotateRight();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 2 times %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 2 times flip cols %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipRows();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 2 times flip cols and rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 2 times flip rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        // -

                        // -
                        $unmatched_tile->rotateRight();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 3 times %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 3 times flip cols %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipRows();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 3 times flip cols and rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 3 times flip rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        // -

                        // -
                        $unmatched_tile->rotateRight();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 4 times %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 4 times flip cols %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipRows();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 4 times flip cols and rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        $unmatched_tile->flipCols();
                        if ($this->check($unmatched_tile, $grid_tile)) {
                            $this->storeInGrid($grid, $this->check($unmatched_tile, $grid_tile), $unmatched_tile, $y, $x);

                            $this->log(sprintf("Rotate right 4 times flip rows %d with %d", $unmatched_tile->getId(), $grid_tile->getId()));
                            unset($tiles[$unmatched_id]);
                            continue;
                        }
                        // -
                    }
                }
            }
        }

        $this->recursive_ksort($grid);

        /** @var Tile[] $top_row */
        $top_row = reset($grid);
        /** @var Tile[] $bottom_row */
        $bottom_row = end($grid);

        $top_left = reset($top_row)->getId();
        $top_right = end($top_row)->getId();
        $bottom_left = reset($bottom_row)->getId();
        $bottom_right = end($bottom_row)->getId();

        return ($top_left * $top_right * $bottom_left * $bottom_right);
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

    public function storeInGrid(array &$grid, int $direction, Tile $tile, int $y, int $x): void
    {
        switch ($direction) {
            case self::LEFT:
                $grid[$y][$x - 1] = $tile;
                break;
            case self::RIGHT:
                $grid[$y][$x + 1] = $tile;
                break;
            case self::UP:
                $grid[$y - 1][$x] = $tile;
                break;
            case self::DOWN:
                $grid[$y + 1][$x] = $tile;
                break;
        }
    }

    public function check(Tile $unmatchedTile, Tile $tileInGrid): ?int
    {
        // Check with no rotation to see if any edges match
        $tileA = $unmatchedTile->getTile();
        $tileB = $tileInGrid->getTile();

        // Check top row to bottom
        if ($tileA[0] === end($tileB)) {
            return self::DOWN;
        }

        // Check bottom row to top
        if (end($tileA) === $tileB[0]) {
            return self::UP;
        }

        $count = count($tileA);
        $strlen = count($tileA[0]) - 1;

        // Check left row matches to right
        $match = 0;
        for ($i = 0; $i < $count; $i++) {
            if ($tileA[$i][0] === $tileB[$i][$strlen]) {
                $match++;
            }
        }
        if ($match === $count) {
            return self::RIGHT;
        }

        // Check right row matches to left
        $count = count($tileA);
        $match = 0;
        for ($i = 0; $i < $count; $i++) {
            if ($tileA[$i][$strlen] === $tileB[$i][0]) {
                $match++;
            }
        }
        if ($match === $count) {
            return self::LEFT;
        }

        return null;
    }

    public function part2(): int|string
    {
        $input = $this->getInput(2020, 20, 2);

        return 0;
    }
}

class Tile
{
    public int $id;
    public array $tile;

    /**
     * Tile constructor.
     * @param int $id
     * @param array $tile
     */
    public function __construct(int $id, array $tile)
    {
        $this->id = $id;
        $this->tile = $tile;
    }

    public function getId(): int
    {
        return $this->id;
    }

    public function setId(int $id): void
    {
        $this->id = $id;
    }

    public function getTile(): array
    {
        return $this->tile;
    }

    public function setTile(array $tile): void
    {
        $this->tile = $tile;
    }

    public function flipRows(): void
    {
        $this->setTile(array_reverse($this->getTile()));
    }

    public function flipCols(): void
    {
        $new = [];
        $tile = $this->getTile();

        foreach ($tile as $row) {
            $new[] = array_reverse($row);
        }
        $this->setTile($new);
    }

    public function rotateRight(): void
    {
        $new = [];
        $tile = $this->getTile();
        $len = count($tile);

        for ($y = 0; $y < $len; $y++) {
            for ($x = 0; $x < $len; $x++) {
                $new[$y][$x] = $tile[$len - 1 - $x][$y];
            }
        }
        $this->setTile($new);
    }
}
