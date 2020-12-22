<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day22 extends Day
{

    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 22, 1);
        unset($input[0]);

        $player1 = [];
        $player2 = [];
        $player2_parse = false;

        foreach ($input as $line) {
            if (empty($line)) {
                continue;
            }
            if ($line === "Player 2:") {
                $player2_parse = true;
                continue;
            }
            if (!$player2_parse) {
                $player1[] = (int)$line;
            } else {
                $player2[] = (int)$line;
            }
        }

        $round = 1;
        while (count($player1) > 0 && count($player2) > 0) {
            $this->log("-- Round " . $round . "--");
            $this->log("Player 1's deck: " . implode(",", $player1));
            $this->log("Player 2's deck: " . implode(",", $player2));

            $p1_card = array_shift($player1);
            $p2_card = array_shift($player2);
            $this->log("Player 1 plays: " . $p1_card);
            $this->log("Player 2 plays: " . $p2_card);

            if ($p1_card > $p2_card) {
                $player1[] = $p1_card;
                $player1[] = $p2_card;
                $this->log("Player 1 wins the round!");
            } elseif ($p2_card > $p1_card) {
                $player2[] = $p2_card;
                $player2[] = $p1_card;
                $this->log("Player 2 wins the round!");
            } else {
                $this->log("equal " . $p1_card);
            }
            $round++;
        }

        $winner = [];
        if (count($player1) > 0) {
            $winner = $player1;
        } else {
            $winner = $player2;
        }

        $this->log("== Post-game results ==");
        $this->log("Player 1's deck: " . implode(",", $player1));
        $this->log("Player 2's deck: " . implode(",", $player2));

        $winner = array_reverse($winner);
        $sum = 0;
        foreach ($winner as $value => $card) {
            $sum += (($value + 1) * $card);
        }

        return $sum;
    }

    public function part2(): int|string
    {
        return 0;

    }

}
