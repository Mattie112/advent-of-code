<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day25Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day25();
    }

    public function dataProvider(): array
    {
        return [
            "day 25 part 1 test" => [14897079, self::PART1, true],
            "day 25 part 1 prod" => [7269858, self::PART1, false],
            "day 25 part 2 test" => [0, self::PART2, true],
            "day 25 part 2 prod" => [0, self::PART2, false],
        ];
    }

    public function part1Provider(): array
    {
        return [
            [17807724, 14897079],
        ];
    }

    /**
     * @dataProvider part1Provider
     * @param int $moves
     * @param int $expected
     */
    public function testDay25Part1Only10Moves(int $moves, int $expected): void
    {
        $day = new Day25();
        $day->setTest(true);
        self::assertSame($expected, $day->part1());
    }

}
