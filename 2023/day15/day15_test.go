package main

import "testing"

func Test_day15Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day15-test.txt"}, 1320},
		{"part 1 answer", args{path: "day15.txt"}, 512797},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day15Part1(tt.args.path); got != tt.want {
				t.Errorf("day15Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day15Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day15-test.txt"}, 145},
		{"part 2 answer", args{path: "day15.txt"}, 262454},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day15Part2(tt.args.path); got != tt.want {
				t.Errorf("day15Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
