<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day16Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day16();
    }

    public function dataProvider(): array
    {
        return [
            "day 16 part 1 test" => [71, self::PART1, true],
            "day 16 part 1 prod" => [20058, self::PART1, false],
            "day 16 part 2 test" => [0, self::PART2, true],
            "day 16 part 2 prod" => [0, self::PART2, false],
        ];
    }

}
