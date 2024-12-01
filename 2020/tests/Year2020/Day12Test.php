<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day12Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day12();
    }

    public function dataProvider(): array
    {
        return [
            "day 12 part 1 test" => [25, self::PART1, true],
            "day 12 part 1 prod" => [820, self::PART1, false],
            "day 12 part 2 test" => [286, self::PART2, true],
            "day 12 part 2 prod" => [66614, self::PART2, false],
        ];
    }

}
