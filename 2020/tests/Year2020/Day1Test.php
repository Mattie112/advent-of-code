<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day1Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day1();
    }

    public function dataProvider(): array
    {
        return [
            "day 1 part 1 test" => [514579, self::PART1, true],
            "day 1 part 1 prod" => [138379, self::PART1, false],
            "day 1 part 2 test" => [241861950, self::PART2, true],
            "day 1 part 2 prod" => [85491920, self::PART2, false],
        ];
    }
}
