<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day12 extends Day
{
    public const NORTH = "N";
    public const SOUTH = "S";
    public const EAST = "E";
    public const WEST = "W";
    public const FORWARD = "F";
    public const ROTATE_RIGHT = "R";
    public const ROTATE_LEFT = "L";

    /**
     * Part 1 and 2 are roughly the same, however part 1 is based on moving the ship and part 2 is based on moving the waypoint
     *
     * @return int|string
     */
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 12, 1);

        $ship_direction = 90; // EAST
        $ship_position_x = 0; // EAST = pos, WEST = neg
        $ship_position_y = 0; // NORTH = NEG, SOUTH = POS

        foreach ($input as $line) {
            preg_match("@([A-Z]+)(\d+)@", $line, $matches);
            $instruction = $matches[1];
            $amount = (int)$matches[2];

            // If we need to go forward simply change the instruction from F to the correct direction and then go on with the regular moving code
            if ($instruction === self::FORWARD) {
                $instruction = match ($ship_direction) {
                    0, 360 => self::NORTH,
                    90 => self::EAST,
                    180 => self::SOUTH,
                    270 => self::WEST,
                };
            }
            switch ($instruction) {
                case self::NORTH:
                    $ship_position_y -= $amount;
                    break;
                case self::SOUTH:
                    $ship_position_y += $amount;
                    break;
                case self::EAST:
                    $ship_position_x += $amount;
                    break;
                case self::WEST:
                    $ship_position_x -= $amount;
                    break;
                case self::ROTATE_RIGHT:
                    $ship_direction = ($ship_direction + $amount) % 360;
                    break;
                case self::ROTATE_LEFT:
                    $ship_direction = ($ship_direction - $amount) % 360;
                    break;
            }
            if ($ship_direction < 0) {
                $ship_direction += 360;
            }
        }

        return abs($ship_position_x) + abs($ship_position_y);
    }

    /**
     * Part 1 and 2 are roughly the same, however part 1 is based on moving the ship and part 2 is based on moving the waypoint
     *
     * @return int|string
     */
    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 12, 1);

        $ship_position_x = 0; // EAST = pos, WEST = neg
        $ship_position_y = 0; // NORTH = NEG, SOUTH = POS
        $waypoint_x = 10;
        $waypoint_y = -1;

        foreach ($input as $line) {
            preg_match("@([A-Z]+)(\d+)@", $line, $matches);
            $instruction = $matches[1];
            $amount = (int)$matches[2];

            // If we need to go forward simply change the instruction from F to the correct direction and then go on with the regular moving code
            if ($instruction === self::FORWARD) {
                $ship_position_x += ($waypoint_x * $amount);
                $ship_position_y += ($waypoint_y * $amount);
            }

            switch ($instruction) {
                case self::NORTH:
                    $waypoint_y -= $amount;
                    break;
                case self::SOUTH:
                    $waypoint_y += $amount;
                    break;
                case self::EAST:
                    $waypoint_x += $amount;
                    break;
                case self::WEST:
                    $waypoint_x -= $amount;
                    break;
                case self::ROTATE_RIGHT:
                    [$waypoint_x, $waypoint_y] = $this->rotateWaypointRelativeToShip($waypoint_x, $waypoint_y, $amount, "R");
                    break;
                case self::ROTATE_LEFT:
                    [$waypoint_x, $waypoint_y] = $this->rotateWaypointRelativeToShip($waypoint_x, $waypoint_y, $amount, "L");
                    break;
            }

            if ($this->output?->isVerbose()) {
                $this->log($line);
                echo "Ship: ";
                $this->debug($ship_position_x, $ship_position_y);
                echo PHP_EOL;
                echo "Wayp: ";
                $this->debug($waypoint_x, $waypoint_y);
                echo PHP_EOL;
                echo PHP_EOL;
            }
        }

        return abs($ship_position_x) + abs($ship_position_y);
    }

    public function rotateWaypointRelativeToShip($waypoint_x, $waypoint_y, $deg, $dir): array
    {
        $rotations = $deg / 90;
        for ($i = 0; $i < $rotations; $i++) {
            $tmpx = $waypoint_x;
            $tmpy = $waypoint_y;
            if ($dir === "L") {
                $waypoint_x = $tmpy;
                $waypoint_y = -$tmpx;
            } else {
                $waypoint_x = -$tmpy;
                $waypoint_y = $tmpx;
            }
        }
        return [$waypoint_x, $waypoint_y];
    }

    public function debug($posx, $posy): void
    {
        if (!$this->output?->isVerbose()) {
            return;
        }
        if ($posx < 0) {
            echo "west " . abs($posx);
        } else {
            echo "east " . abs($posx);
        }
        echo "  ";
        if ($posy < 0) {
            echo "north " . abs($posy);
        } else {
            echo "south " . abs($posy);
        }
    }
}
