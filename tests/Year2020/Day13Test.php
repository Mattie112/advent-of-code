<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day13Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day13();
    }

    public function dataProvider(): array
    {
        return [
            "day 13 part 1 test" => [295, self::PART1, true],
            "day 13 part 1 prod" => [6568, self::PART1, false],
            "day 13 part 2 test" => [1068781, self::PART2, true],
            "day 13 part 2 prod" => [554865447501099, self::PART2, false],
        ];
    }

    public function moreInputProvider(): array
    {
        return [
            ["a\n17,x,13,19", 3417],
            ["a\n67,7,59,61", 754018],
            ["a\n67,x,7,59,61", 779210],
            ["a\n67,7,x,59,61", 1261476],
            ["a\n1789,37,47,1889", 1202161486],
        ];
    }

    /**
     * @dataProvider moreInputProvider
     * @param string $test_input
     * @param int $expected
     */
    public function testDay13AlternateTestInput(string $test_input, int $expected): void
    {
        /** @var MockObject|Day13 $day */
        $day = $this->getMockBuilder(Day13::class)->onlyMethods(["getInput"])->getMock();
        $day->method("getInput")->willReturn($test_input);
        self::assertSame($expected, $day->part2());
    }

}
