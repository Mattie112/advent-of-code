package main

import "testing"

func Test_day3Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day3-test.txt"}, 198},
		{"part 1 answer", args{path: "day3.txt"}, 2972336},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2Part1(tt.args.path); got != tt.want {
				t.Errorf("day2Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day3Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day3-test.txt"}, 230},
		{"part 2 answer", args{path: "day3.txt"}, 3368358},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2Part2(tt.args.path); got != tt.want {
				t.Errorf("day2Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
