package main

import "testing"

func Test_day11Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day11-test.txt"}, 55312},
		{"part 1 answer", args{path: "day11.txt"}, 202019},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.path, 25); got != tt.want {
				t.Errorf("day11Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day11Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day11-test.txt"}, 65601038650482},
		{"part 2 answer", args{path: "day11.txt"}, 239321955280205},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.path, 75); got != tt.want {
				t.Errorf("day11Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
