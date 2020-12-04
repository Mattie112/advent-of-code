<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day3Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day3();
    }

    public function dataProvider(): array
    {
        return [
            "day 3 part 1 test" => [7, self::PART1, true],
            "day 3 part 1 prod" => [276, self::PART1, false],
            "day 3 part 2 test" => [336, self::PART2, true],
            "day 3 part 2 prod" => [7812180000, self::PART2, false],
        ];
    }
}
