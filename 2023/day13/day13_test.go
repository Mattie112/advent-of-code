package main

import "testing"

func Test_day13Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day13-test.txt"}, 405},
		{"part 1 answer", args{path: "day13.txt"}, 30802},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day13Part1(tt.args.path); got != tt.want {
				t.Errorf("day13Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day13Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day13-test.txt"}, 400},
		{"part 2 answer", args{path: "day13.txt"}, 37876},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day13Part2(tt.args.path); got != tt.want {
				t.Errorf("day13Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
