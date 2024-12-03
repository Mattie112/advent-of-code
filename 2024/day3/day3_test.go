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
		{"part 1 test", args{path: "day3-test.txt"}, 161},
		{"part 1 answer", args{path: "day3.txt"}, 161289189},
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
		{"part 2 test", args{path: "day3-test-part2.txt"}, 48},
		{"part 2 answer", args{path: "day3.txt"}, 83595109},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1Part2(tt.args.path); got != tt.want {
				t.Errorf("day1Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
