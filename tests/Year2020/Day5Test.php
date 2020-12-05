<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day5Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day5();
    }

    public function dataProvider(): array
    {
        return [
            "day 5 part 1 test" => [820, self::PART1, true],
            "day 5 part 1 prod" => [858, self::PART1, false],
            "day 5 part 2 test" => [17, self::PART2, true],
            "day 5 part 2 prod" => [557, self::PART2, false],
        ];
    }
}
