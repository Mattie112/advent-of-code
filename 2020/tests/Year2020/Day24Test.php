<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day24Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day24();
    }

    public function dataProvider(): array
    {
        return [
            "day 24 part 1 test" => [10, self::PART1, true],
            "day 24 part 1 prod" => [351, self::PART1, false],
            "day 24 part 2 test" => [2208, self::PART2, true],
            "day 24 part 2 prod" => [3869, self::PART2, false],
        ];
    }
}
