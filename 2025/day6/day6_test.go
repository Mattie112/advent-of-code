package main

import "testing"

func Test_day6Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day6-test.txt"}, 4277556},
		{"part 1 answer", args{path: "day6.txt"}, 3785892992137},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day6Part1(tt.args.path); got != tt.want {
				t.Errorf("day6Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day6Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day6-test.txt"}, 3263827},
		{"part 2 answer", args{path: "day6.txt"}, 7669802156452},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day6Part2(tt.args.path); got != tt.want {
				t.Errorf("day6Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
