package main

import "testing"

func Test_day1Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day1-test.txt"}, 7},
		{"part 1 answer", args{path: "day1.txt"}, 1292},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1Part1(tt.args.path); got != tt.want {
				t.Errorf("day1Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day1Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day1-test.txt"}, 5},
		{"part 2 answer", args{path: "day1.txt"}, 1262},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1Part2(tt.args.path); got != tt.want {
				t.Errorf("day1Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
