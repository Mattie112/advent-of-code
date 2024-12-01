<?php

namespace mattie112\AdventOfCode\Year2020;

use JetBrains\PhpStorm\ArrayShape;
use mattie112\AdventOfCode\BaseTest;

class Day9Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day9();
    }

    #[ArrayShape(["day 9 part 1 test" => "array", "day 9 part 1 prod" => "array", "day 9 part 2 test" => "array", "day 9 part 2 prod" => "array"])] public function dataProvider(): array
    {
        return [
            "day 9 part 1 test" => [127, self::PART1, true],
            "day 9 part 1 prod" => [15690279, self::PART1, false],
            "day 9 part 2 test" => [62, self::PART2, true],
            "day 9 part 2 prod" => [2174232, self::PART2, false],
        ];
    }
}
