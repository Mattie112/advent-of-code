package main

import "testing"

func Test_day5Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day5-test.txt"}, 2},
		{"part 1 answer", args{path: "day5.txt"}, 580},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day5Part1(tt.args.path); got != tt.want {
				t.Errorf("day5Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day5Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day5-test.txt"}, 4},
		{"part 2 answer", args{path: "day5.txt"}, 895},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day5Part2(tt.args.path); got != tt.want {
				t.Errorf("day5Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
