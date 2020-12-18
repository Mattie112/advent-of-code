<?php /** @noinspection DuplicatedCode */

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day16 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 16, 1, PHP_EOL . PHP_EOL);

        // Get the 2 things I need
        $fields_input = explode(PHP_EOL, $input[0]);
        $nearby_tickets_input = explode(PHP_EOL, $input[2]);
        unset($nearby_tickets_input[0]); // Drop the "Nearby tickets" line

        $fields = $this->getFields($fields_input);

        $nearby_tickets = $this->getNearbyTickets($nearby_tickets_input);

        return $this->getValidTickets($nearby_tickets, $fields)[0];
    }

    public function getNearbyTickets($nearby_tickets_input): array
    {
        $nearby_tickets = [];
        foreach ($nearby_tickets_input as $n) {
            preg_match_all("/(\d+),?/", $n, $matches);
            $nearby_tickets[] = array_map('intval', $matches[1]);
        }
        return $nearby_tickets;
    }

    public function getFields($fields_input): array
    {
        $fields = [];
        foreach ($fields_input as $f) {
            preg_match("/([a-z]+ ?[a-z]*): (\d+)-(\d+) or (\d+)-(\d+)/", $f, $matches);
            [, $field, $min_a, $max_a, $min_b, $max_b] = $matches;
            $fields[$field] = [[(int)$min_a, (int)$max_a], [(int)$min_b, (int)$max_b]];
        }
        return $fields;
    }

    public function getValidTickets(array $nearby_tickets, array $fields): array
    {
        $valid_tickets = $nearby_tickets;
        $invalid_values = [];
        foreach ($nearby_tickets as $ticket_id => $ticket) {
            $fields_valid = 0;
            foreach ($ticket as $i => $ticket_value) {
                $values_valid = 0;
                foreach ($fields as $field) {
                    if (($ticket_value >= $field[0][0] && $ticket_value <= $field[0][1]) || ($ticket_value >= $field[1][0] && $ticket_value <= $field[1][1])) {
                        $values_valid++;
                    }
                }
                if ($values_valid === 0) {
                    $invalid_values[] = $ticket_value;
                    unset($valid_tickets[$ticket_id]);
                }
            }

        }
        return [array_sum($invalid_values), array_values($valid_tickets)];
    }

    public function part2(): int|string
    {
        ini_set("memory_limit", "-1");
        $input = $this->getInputAsArray(2020, 16, 2, PHP_EOL . PHP_EOL);

        // Get the 3 things I need
        $fields_input = explode(PHP_EOL, $input[0]);
        $my_ticket_input = explode(PHP_EOL, $input[1])[1];
        $nearby_tickets_input = explode(PHP_EOL, $input[2]);
        unset($nearby_tickets_input[0]); // Drop the "Nearby tickets" line

        $fields = $this->getFields($fields_input);

        $nearby_tickets = $this->getNearbyTickets($nearby_tickets_input);

        $valid_tickets = $this->getValidTickets($nearby_tickets, $fields)[1];
        $ticket_field_count = count($valid_tickets[0]);

        // As we need to find out the fields we are looping by field (and not by ticket)
        $found_fields = [];
        $remaining_fields = $fields;
        $ignore_i = []; // The field (or value) indexes that have already been found
        while (count($remaining_fields) > 0) {
            $possible_found_fields = [];
            $tmp_remaining_fields = $remaining_fields;
            $tmp_ignore_i = $ignore_i; // The field (or value) indexes that have already been found

            // todo dont loop through te fields, loop through $i as any column dan match to multiple fields, if we have only 1 match we know our field for certain
            // so sad refactor :(
            foreach ($tmp_remaining_fields as $field_id => $field_ranges) {

                for ($i = 0; $i < $ticket_field_count; $i++) {
                    if (isset($tmp_ignore_i[$i])) {
                        // If we have alredy found a field for this field we can skip it
                        continue;
                    }

                    $valid = 0;
                    foreach ($valid_tickets as $ticket) {
                        $possible_field_value = $ticket[$i];
                        if (($possible_field_value >= $field_ranges[0][0] && $possible_field_value <= $field_ranges[0][1]) || ($possible_field_value >= $field_ranges[1][0] && $possible_field_value <= $field_ranges[1][1])) {
                            $valid++;
                        } else {
                            continue 2;
                        }
                    }
                    if ($valid === count($valid_tickets)) {
//                        $possible_found_fields[$i][] = $field_id;
                        $possible_found_fields[$field_id][] = $i;
                        $tmp_ignore_i[$i] = true;
                        unset($tmp_remaining_fields[$field_id]);

//                         Now we are sure on this field
//                        $found_fields[$i] = $field_id;
                        $this->log("Found possible field: " . $field_id . " at position: " . $i);
//                        continue 2; // go to next field
                    }
                }
                $a = 1;

            }


            // Now extract the found thingies
            foreach ($possible_found_fields as $field_name => $field) {
                if (count($field) === 1) {
                    $position = $field[0];
                    // Only one hit for this fit we can be sure we have the correct one
                    unset($remaining_fields[$field_name]);
                    // $field[0] == position
                    $ignore_i[$position] = true;
                    $found_fields[$field_name] = $position; // or $field_id??
                    $this->log("");
                    $this->log("I am now certain on field " . $field_name . " at position: " . $position);
                    $this->log("");
                }
            }

            $this->log("Went through everything, starting over");
        }

        $my_ticket = explode(",", $my_ticket_input);
        $mul = [];
        foreach ($found_fields as $id => $field) {
            $this->log($field . " = " . $my_ticket[$id]);
            if (str_starts_with($field, "departure")) {
                $mul[] = $my_ticket[$id];
            }
        }

        return array_product($mul);
    }


}
