<?php

namespace mattie112\AdventOfCode\Year2020;

use JetBrains\PhpStorm\ArrayShape;
use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day8Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day8();
    }

    #[ArrayShape(["day 8 part 1 test" => "array", "day 8 part 1 prod" => "array", "day 8 part 2 test" => "array", "day 8 part 2 prod" => "array"])] public function dataProvider(): array
    {
        return [
            "day 8 part 1 test" => [5, self::PART1, true],
            "day 8 part 1 prod" => [1521, self::PART1, false],
            "day 8 part 2 test" => [8, self::PART2, true],
            "day 8 part 2 prod" => [1016, self::PART2, false],
        ];
    }
}
