<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day17Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day17();
    }

    public function dataProvider(): array
    {
        return [
            "day 17 part 1 test" => [112, self::PART1, true],
            "day 17 part 1 prod" => [388, self::PART1, false],
            "day 17 part 2 test" => [848, self::PART2, true],
            "day 17 part 2 prod" => [2280, self::PART2, false],
        ];
    }

}
