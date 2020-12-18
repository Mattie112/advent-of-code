<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day18Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day18();
    }

    public function dataProvider(): array
    {
        return [
            "day 18 part 1 test" => [71, self::PART1, true],
            "day 18 part 1 prod" => [5374004645253, self::PART1, false],
            "day 18 part 2 test" => [0, self::PART2, true],
            "day 18 part 2 prod" => [0, self::PART2, false],
        ];
    }

    public function moreInputProvider(): array
    {
        return [
            ["1 + (2 * 3) + (4 * (5 + 6))", 51],
            ["2 * 3 + (4 * 5)", 26],
            ["5 + (8 * 3 + 9 + 3 * 4 * 3)", 437],
            ["5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240],
            ["((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632],
        ];
    }

    /**
     * @dataProvider moreInputProvider
     * @param string $test_input
     * @param int $expected
     */
    public function testDay18Part1AlternateTestInput(string $test_input, int $expected): void
    {
        /** @var MockObject|Day18 $day */
        $day = $this->getMockBuilder(Day18::class)->onlyMethods(["getInput"])->getMock();
        $day->method("getInput")->willReturn($test_input);
        self::assertSame($expected, $day->part1());
    }
}
