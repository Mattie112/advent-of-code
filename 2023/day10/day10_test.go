package main

import "testing"

func Test_day10Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day10-test-1.txt"}, 4},
		{"part 1 test", args{path: "day10-test-2.txt"}, 4},
		{"part 1 test", args{path: "day10-test-3.txt"}, 8},
		{"part 1 test", args{path: "day10-test-4.txt"}, 8},
		{"part 1 answer", args{path: "day10.txt"}, 6870},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day10Part1(tt.args.path); got != tt.want {
				t.Errorf("day10Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day10Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day10-test-1.txt"}, -1},
		{"part 2 answer", args{path: "day10.txt"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day10Part2(tt.args.path); got != tt.want {
				t.Errorf("day10Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
