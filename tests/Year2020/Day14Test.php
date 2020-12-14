<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day14Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day14();
    }

    public function dataProvider(): array
    {
        return [
            "day 14 part 1 test" => [165, self::PART1, true],
            "day 14 part 1 prod" => [11926135976176, self::PART1, false],
            "day 14 part 2 test" => [208, self::PART2, true],
            "day 14 part 2 prod" => [4330547254348, self::PART2, false],
        ];
    }

}
