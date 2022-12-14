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
		{"part 1 test", args{path: "day3-test.txt"}, 157},
		{"part 1 answer", args{path: "day3.txt"}, 7795},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day3Part1(tt.args.path); got != tt.want {
				t.Errorf("day3Part1() = %v, want %v", got, tt.want)
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
		{"part 2 test", args{path: "day3-test.txt"}, 70},
		{"part 2 answer", args{path: "day3.txt"}, 2703},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day3Part2(tt.args.path); got != tt.want {
				t.Errorf("day3Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
