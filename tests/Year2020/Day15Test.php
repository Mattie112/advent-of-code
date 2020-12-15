<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day15Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day15();
    }

    public function dataProvider(): array
    {
        return [
            "day 15 part 1 test" => [436, self::PART1, true],
            "day 15 part 1 prod" => [517, self::PART1, false],
            "day 15 part 2 test" => [0, self::PART2, true],
            "day 15 part 2 prod" => [0, self::PART2, false],
        ];
    }

    public function moreInputProvider(): array
    {
        return [
            ["1,3,2", 1],
            ["2,1,3", 10],
            ["1,2,3", 27],
            ["2,3,1", 78],
            ["3,2,1", 438],
            ["3,1,2", 1836],
        ];
    }

    /**
     * @dataProvider moreInputProvider
     * @param string $test_input
     * @param int $expected
     */
    public function testDay15Part1AlternateTestInput(string $test_input, int $expected): void
    {
        /** @var MockObject|Day15 $day */
        $day = $this->getMockBuilder(Day15::class)->onlyMethods(["getInput"])->getMock();
        $day->method("getInput")->willReturn($test_input);
        self::assertSame($expected, $day->part1());
    }
}
