<?php

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\BaseTest;
use PHPUnit\Framework\MockObject\MockObject;

class Day7Test extends BaseTest
{
    public function setUp(): void
    {
        $this->day = new Day7();
    }

    public function dataProvider(): array
    {
        return [
            "day 7 part 1 test" => [4, self::PART1, true],
            "day 7 part 1 prod" => [296, self::PART1, false],
            "day 7 part 2 test" => [32, self::PART2, true],
            "day 7 part 2 prod" => [9339, self::PART2, false],
        ];
    }

    public function testDay7Part2AlternateTestInput(): void
    {
        $test_input = <<<EOT
shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
EOT;

        /** @var MockObject|Day7 $day */
        $day = $this->getMockBuilder(Day7::class)->onlyMethods(["getInput"])->getMock();
        $day->expects(self::once())->method("getInput")->willReturn($test_input);
        self::assertSame(126, $day->part2());
    }
}
