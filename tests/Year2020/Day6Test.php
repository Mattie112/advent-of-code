<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day6Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day6();
    }

    public function dataProvider(): array
    {
        return [
            "day 6 part 1 test" => [11, self::PART1, true],
            "day 6 part 1 prod" => [6680, self::PART1, false],
            "day 6 part 2 test" => [6, self::PART2, true],
            "day 6 part 2 prod" => [3117, self::PART2, false],
        ];
    }

    /**
     * Just test for my alternate solution
     */
    public function testMaps(): void
    {
        /** @var Day6 $day */
        $day = $this->day;
        $day->setTest(true);
        self::assertEquals(6, $day->part2WithMaps());
        $day->setTest(false);
        self::assertEquals(3117, $day->part2WithMaps());
    }
}
