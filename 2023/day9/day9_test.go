package main

import "testing"

func Test_day9Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day9-test.txt"}, 114},
		{"part 1 answer", args{path: "day9.txt"}, 2008960228},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day9Part1(tt.args.path); got != tt.want {
				t.Errorf("day9Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day9Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day9-test.txt"}, 2},
		{"part 2 answer", args{path: "day9.txt"}, 1097},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day9Part2(tt.args.path); got != tt.want {
				t.Errorf("day9Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
