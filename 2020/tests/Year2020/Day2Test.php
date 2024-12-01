<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day2Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day2();
    }

    public function dataProvider(): array
    {
        return [
            "day 2 part 1 test" => [2, self::PART1, true],
            "day 2 part 1 prod" => [580, self::PART1, false],
            "day 2 part 2 test" => [1, self::PART2, true],
            "day 2 part 2 prod" => [611, self::PART2, false],
        ];
    }
}
