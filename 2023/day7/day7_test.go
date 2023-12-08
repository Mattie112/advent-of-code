package main

import "testing"

func Test_day7Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day7-test.txt"}, 6440},
		{"part 1 answer", args{path: "day7.txt"}, 252656917},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day7Part1(tt.args.path); got != tt.want {
				t.Errorf("day7Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day7Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day7-test.txt"}, 5905},
		{"part 2 answer", args{path: "day7.txt"}, 253499763},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day7Part2(tt.args.path); got != tt.want {
				t.Errorf("day7Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
