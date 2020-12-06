<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day6 extends Day
{

    public function part1(): int|string
    {
        $input = $this->getInput(2020, 6, 1);
        $input = explode("\n\n", $input);

        $total_answers = 0;
        foreach ($input as $group) {
            $persons = explode("\n", $group);
            $answers = [];
            foreach ($persons as $person) {
                $person_answers = str_split($person);
                foreach ($person_answers as $answer) {
                    isset($answers[$answer]) ? $answers[$answer]++ : $answers[$answer] = 1;
                }
            }
            $total_answers += count($answers);
        }

        return $total_answers;
    }


    public function part2(): int|string
    {
        $input = $this->getInput(2020, 6, 1);
        $input = explode("\n\n", $input);

        $total_answers = 0;
        foreach ($input as $group) {
            $persons = explode("\n", $group);
            $person_count = count($persons);
            $answers = [];
            foreach ($persons as $person) {
                $person_answers = str_split($person);

                foreach ($person_answers as $answer) {
                    isset($answers[$answer]) ? $answers[$answer]++ : $answers[$answer] = 1;
                }
            }

            // If we have more persons only take the answers with the amount equal to the amount of persons
            $total_answers += count(array_filter($answers, static function ($elem) use ($person_count) {
                return $elem === $person_count;
            }));
        }

        return $total_answers;
    }

    /**
     * This is basically the same as part 1 but then with some array_maps but as it really makes it more unreadable I just use the original method
     *
     * @return int|string
     */
    public function part2WithMaps(): int|string
    {
        $input = $this->getInput(2020, 6, 1);
        $input = explode("\n\n", $input);

        $total_answers = 0;
        foreach ($input as $group) {
            $persons = explode("\n", $group);
            $person_count = count($persons);
            $answers = [];

            $splitted = array_map(static function ($elem) {
                return str_split($elem);
            }, $persons);

            array_map(static function ($person_answers) use (&$answers) {
                array_map(static function ($answer) use (&$answers) {
                    isset($answers[$answer]) ? $answers[$answer]++ : $answers[$answer] = 1;
                }, $person_answers);
            }, $splitted);

            // If we have more persons only take the answers with the amount equal to the amount of persons
            $total_answers += count(array_filter($answers, static function ($elem) use ($person_count) {
                return $elem === $person_count;
            }));
        }

        return $total_answers;
    }
}
