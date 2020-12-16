<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day16 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 16, 1, PHP_EOL . PHP_EOL);

        // Get the 3 things I need
        $fields_input = explode(PHP_EOL, $input[0]);
        $my_ticket_input = explode(PHP_EOL, $input[1])[1];
        $nearby_tickets_input = explode(PHP_EOL, $input[2]);
        unset($nearby_tickets_input[0]);

        $fields = [];
        foreach ($fields_input as $f) {
            preg_match("/([a-z]+): (\d+)-(\d+) or (\d+)-(\d+)/", $f, $matches);
            [, $field, $min_a, $max_a, $min_b, $max_b] = $matches;
            $fields[$field] = [[(int)$min_a, (int)$max_a], [(int)$min_b, (int)$max_b]];
        }

        $nearby_tickets = [];
        foreach ($nearby_tickets_input as $n) {
            preg_match_all("/(\d+),?/", $n, $matches);
            $nearby_tickets[] = array_map('intval', $matches[1]);
        }

        // Now we have to check if the ticket is 100% invalid
        $valid_tickets = [];
        $invalid_tickets = [];
        $invalid_values = [];
        foreach ($nearby_tickets as $ticket) {
            $fields_valid = 0;
            foreach ($ticket as $ticket_field) {
                $values_valid = 0;
                foreach ($fields as $field) {
                    if (($ticket_field >= $field[0][0] && $ticket_field <= $field[0][1]) || ($ticket_field >= $field[1][0] && $ticket_field <= $field[1][1])) {
                        $fields_valid++;
                        $values_valid++;
                    }
                }
                if ($values_valid === 0) {
                    $invalid_values[] = $ticket_field;
                }
            }
            if ($fields_valid > 0) {
                $valid_tickets[] = $ticket;
            } else {
                $invalid_tickets[] = $ticket;
            }
        }

        return array_sum($invalid_values);
    }


    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 15, 2);

        return 0;
    }


}
