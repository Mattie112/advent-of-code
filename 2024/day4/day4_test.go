package main

import "testing"

func Test_day4Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day4-test.txt"}, 18},
		{"part 1 answer", args{path: "day4.txt"}, 2401},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.path); got != tt.want {
				t.Errorf("day4Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day4Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day4-test.txt"}, 9},
		{"part 2 answer", args{path: "day4.txt"}, 1822},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.path); got != tt.want {
				t.Errorf("day4Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
