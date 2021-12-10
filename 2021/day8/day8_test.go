package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day8-test.txt"}, 0},
		{"part 1 test", args{path: "day8-test2.txt"}, 26},
		{"part 1 answer", args{path: "day8.txt"}, 261},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.path); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day8-test.txt"}, 5353},
		{"part 2 test", args{path: "day8-test2.txt"}, 31229},
		{"part 2 answer", args{path: "day8.txt"}, 987553},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.path); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
