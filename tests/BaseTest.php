<?php

namespace mattie112\AdventOfCode;

use PHPUnit\Framework\TestCase;

abstract class BaseTest extends TestCase
{
    public Day $day;
    public const PART1 = 1;
    public const PART2 = 2;

    /**
     * @dataProvider dataProvider
     * @param mixed $expected
     * @param int $part
     * @param bool $test
     */
    public function testDay(mixed $expected, int $part, bool $test): void
    {
        $this->day->setTest($test);
        self::assertSame($expected, $this->day->{"part" . $part}());
    }

    abstract public function dataProvider(): array;
}
