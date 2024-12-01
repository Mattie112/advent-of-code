<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day10Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day10();
    }

    public function dataProvider(): array
    {
        return [
            "day 10 part 1 test" => [220, self::PART1, true],
            "day 10 part 1 prod" => [2112, self::PART1, false],
            "day 10 part 2 test" => [19208, self::PART2, true],
            "day 10 part 2 prod" => [3022415986688, self::PART2, false],
        ];
    }

    public function testDay10AlternateTestInput(): void
    {
        $test_input = <<<EOT
16
10
15
5
1
11
7
19
6
12
4
EOT;

        /** @var MockObject|Day10 $day */
        $day = $this->getMockBuilder(Day10::class)->onlyMethods(["getInput"])->getMock();
        $day->expects(self::any())->method("getInput")->willReturn($test_input);
        self::assertSame(35, $day->part1());
        self::assertSame(8, $day->part2());
    }
}
