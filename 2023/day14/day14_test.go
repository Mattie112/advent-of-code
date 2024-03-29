package main

import "testing"

func Test_day14Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day14-test.txt"}, 136},
		{"part 1 answer", args{path: "day14.txt"}, 109755},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day14Part1(tt.args.path); got != tt.want {
				t.Errorf("day14Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day14Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day14-test.txt"}, 64},
		{"part 2 answer", args{path: "day14.txt"}, 90928},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day14Part2(tt.args.path); got != tt.want {
				t.Errorf("day14Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
