<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;

class Day23Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day23();
    }

    public function dataProvider(): array
    {
        return [
            "day 23 part 1 test" => ["67384529", self::PART1, true],
            "day 23 part 1 prod" => ["52937846", self::PART1, false],
            "day 23 part 2 test" => [0, self::PART2, true],
            "day 23 part 2 prod" => [0, self::PART2, false],
        ];
    }

    public function part1Provider(): array
    {
        return [
            [10, "92658374"],
        ];
    }

    /**
     * @dataProvider part1Provider
     * @param int $moves
     * @param string $expected
     */
    public function testDay23Part1Only10Moves(int $moves, string $expected): void
    {
        $day = new Day23();
        $day->setTest(true);
        $day->moves = $moves;
        self::assertSame($expected, $day->part1());
    }

}
