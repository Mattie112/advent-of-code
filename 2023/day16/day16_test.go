package main

import "testing"

func Test_day16Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day16-test.txt"}, 46},
		{"part 1 answer", args{path: "day16.txt"}, 8539},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day16Part1(tt.args.path); got != tt.want {
				t.Errorf("day16Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day16Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day16-test.txt"}, 51},
		{"part 2 answer", args{path: "day16.txt"}, 8674},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day16Part2(tt.args.path); got != tt.want {
				t.Errorf("day16Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
