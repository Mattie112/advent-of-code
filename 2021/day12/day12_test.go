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
		{"part 1 test", args{path: "day12-test.txt"}, 10},
		{"part 1 test", args{path: "day12-test2.txt"}, 19},
		{"part 1 test", args{path: "day12-test3.txt"}, 226},
		{"part 1 answer", args{path: "day12.txt"}, 3369},
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
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day12-test.txt"}, 36},
		{"part 2 test", args{path: "day12-test2.txt"}, 103},
		{"part 2 test", args{path: "day12-test3.txt"}, 3509},
		{"part 2 answer", args{path: "day12.txt"}, 85883},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.path); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
