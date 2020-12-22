<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day22Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day22();
    }

    public function dataProvider(): array
    {
        return [
            "day 22 part 1 test" => [306, self::PART1, true],
            "day 22 part 1 prod" => [31314, self::PART1, false],
            "day 22 part 2 test" => [291, self::PART2, true],
            "day 22 part 2 prod" => [0, self::PART2, false],
        ];
    }

}
