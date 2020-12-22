<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day19Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day19();
    }

    public function dataProvider(): array
    {
        return [
            "day 19 part 1 test" => [2, self::PART1, true],
            "day 19 part 1 prod" => [160, self::PART1, false],
            "day 19 part 2 test" => [0, self::PART2, true],
            "day 19 part 2 prod" => [0, self::PART2, false],
        ];
    }
}
