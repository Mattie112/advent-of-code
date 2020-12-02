<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode;

use RuntimeException;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

abstract class Day extends Command
{

    public function __construct()
    {
        // I am to lazy to set a name every day ;)
        preg_match("@Day(\d*)@", static::class, $matches);
        $this->setName("day" . $matches[1]);
        Command::__construct();
    }

    abstract public function part1(): int|string;

    abstract public function part2(): int|string;


    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $part1 = $this->part1();
        $output->writeln(static::class . " Part1 answer: " . $part1);
        $part2 = $this->part2();
        $output->writeln(static::class . " Part2 answer: " . $part2);
        return 0;
    }

    public function getInput(int $year, int $day, int $part, bool $test = false): string
    {
        if ($test) {
            if (file_exists(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . "-test.txt")) {
                return trim(file_get_contents(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . ".txt"));
            }
            if (file_exists(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . "-part" . $part . ".txt")) {
                return trim(file_get_contents(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . "-part" . $part . "-test.txt"));
            }
        }

        if (file_exists(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . ".txt")) {
            return trim(file_get_contents(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . ".txt"));
        }
        if (file_exists(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . "-part" . $part . ".txt")) {
            return trim(file_get_contents(__DIR__ . "/../inputs/" . $year . "/day" . $day . "/day" . $day . "-part" . $part . ".txt"));
        }
        throw new RuntimeException("Could not find input");
    }

}
