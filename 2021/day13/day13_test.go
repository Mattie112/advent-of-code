package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day13-test.txt"}, 17},
		{"part 1 answer", args{path: "day13.txt"}, 661},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.path); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		path string
	}

	part1str := "#####.............................................\n#...#.............................................\n#...#.............................................\n#...#.............................................\n#####.............................................\n..................................................\n"
	part2str := "###..####.#..#.#....#..#..##..####.###............\n#..#.#....#.#..#....#.#..#..#.#....#..#...........\n#..#.###..##...#....##...#....###..#..#...........\n###..#....#.#..#....#.#..#....#....###............\n#....#....#.#..#....#.#..#..#.#....#..............\n#....#....#..#.####.#..#..##..#....#..............\n"
	tests := []struct {
		name string
		args args
		want string
	}{
		{"part 2 test", args{path: "day13-test.txt"}, part1str},
		{"part 2 answer", args{path: "day13.txt"}, part2str},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.path); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
