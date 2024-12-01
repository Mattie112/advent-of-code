<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day4Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day4();
    }

    public function dataProvider(): array
    {
        return [
            "day 4 part 1 test" => [2, self::PART1, true],
            "day 4 part 1 prod" => [202, self::PART1, false],
            "day 4 part 2 test" => [4, self::PART2, true],
            "day 4 part 2 prod" => [137, self::PART2, false],
        ];
    }
}
