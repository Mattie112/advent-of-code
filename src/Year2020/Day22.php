<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day22 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 22, 1);

        [$player1, $player2] = $this->getPlayerDecks($input);
        [$player1, $player2] = $this->playNonRecursive($player1, $player2);

        return $this->decideWinner($player1, $player2);
    }

    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 22, 2);

        [$player1, $player2] = $this->getPlayerDecks($input);
        [$player1, $player2] = $this->playRecursive($player1, $player2, 1);

        return $this->decideWinner($player1, $player2);
    }

    /**
     * This is 99% the same as in playRecursive but kept it separate for readability (see comment there)
     *
     * @param array $player1
     * @param array $player2
     * @return array[]
     */
    public function playNonRecursive(array $player1, array $player2): array
    {
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

        return [$player1, $player2];
    }

    public function playRecursive(array $player1, array $player2, int $game_id): array
    {
        $this->log(sprintf("=== Game %d ===", $game_id));
        $previous_for_this_round[1] = [];
        $previous_for_this_round[2] = [];

        $round = 1;
        $winner = 0;
        $winners_card = 0;
        $p1_card_prev = 0;
        while (count($player1) > 0 && count($player2) > 0) {
            $p1_imploded = implode(",", $player1);
            $p2_imploded = implode(",", $player2);

            // If we had these hands before we know we have an infinite loop and just let player 1 win
            if (in_array($p1_imploded, $previous_for_this_round[1], false) || in_array($p2_imploded, $previous_for_this_round[2], false)) {
                return [$player1, $player2, 1, $p1_card_prev];
            }

            $previous_for_this_round[1][] = $p1_imploded;
            $previous_for_this_round[2][] = $p2_imploded;

            $this->log(sprintf("-- Round %d (Game %d) --", $round, $game_id));
            $this->log("Player 1's deck: " . $p1_imploded);
            $this->log("Player 2's deck: " . $p2_imploded);

            // Draw cards
            $p1_card = array_shift($player1);
            $p1_card_prev = $p1_card;
            $p2_card = array_shift($player2);
            $this->log("Player 1 plays: " . $p1_card);
            $this->log("Player 2 plays: " . $p2_card);

            // First determine the amount of cards to see if we need to recurse or not
            if (count($player1) >= $p1_card && count($player2) >= $p2_card) {
                $this->log("Playing a sub-game to determine the winner ...");
                // go recursive
                $copy1 = $player1;
                $copy2 = $player2;
                $tmp_p_1 = array_splice($copy1, 0, $p1_card);
                $tmp_p_2 = array_splice($copy2, 0, $p2_card);
                [, , $winner, $winners_card] = $this->playRecursive($tmp_p_1, $tmp_p_2, $game_id + 1);
                $this->log("");
                $this->log(sprintf("... anyway, back to game %d", $game_id));
                $this->log(sprintf("Player %d wins round %d of game %d!", $winner, $round, $game_id));
                if ($winner === 1) {
                    $player1[] = $p1_card;
                    $player1[] = $p2_card;
                } else {
                    $player2[] = $p2_card;
                    $player2[] = $p1_card;
                }
                $this->log("");
            } else {
                // Just the regular higher then thingy (from playNonRecursive but with winners card)
                /** @noinspection NestedPositiveIfStatementsInspection */
                if ($p1_card > $p2_card) {
                    $player1[] = $p1_card;
                    $player1[] = $p2_card;
                    $winner = 1;
                    $winners_card = $p1_card;
                    $this->log(sprintf("Player 1 wins round %d of game %d!", $round, $game_id));
                } elseif ($p2_card > $p1_card) {
                    $player2[] = $p2_card;
                    $player2[] = $p1_card;
                    $winner = 2;
                    $winners_card = $p2_card;
                    $this->log(sprintf("Player 2 wins round %d of game %d!", $round, $game_id));
                } else {
                    $this->log("Should not happen, equal: " . $p1_card);
                }
            }

            $round++;
        }

        $this->log(sprintf("The winner of game %d is player %d!", $game_id, $winner));
        return [$player1, $player2, $winner, $winners_card];
    }

    public function decideWinner(array $player1, array $player2): float|int
    {
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

    public function getPlayerDecks(array $input): array
    {
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
        return [$player1, $player2];
    }

}
