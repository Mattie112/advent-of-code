<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day21Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day21();
    }

    public function dataProvider(): array
    {
        return [
            "day 21 part 1 test" => [5, self::PART1, true],
            "day 21 part 1 prod" => [2078, self::PART1, false],
            "day 21 part 2 test" => [0, self::PART2, true],
            "day 21 part 2 prod" => [0, self::PART2, false],
        ];
    }

}
