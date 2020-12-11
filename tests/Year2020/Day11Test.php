<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day11Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day11();
    }

    public function dataProvider(): array
    {
        return [
            "day 11 part 1 test" => [37, self::PART1, true],
            "day 11 part 1 prod" => [2418, self::PART1, false],
            "day 11 part 2 test" => [26, self::PART2, true],
            "day 11 part 2 prod" => [2144, self::PART2, false],
        ];
    }


}
