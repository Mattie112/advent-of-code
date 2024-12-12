package main

import "testing"

func Test_day12Part1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day12-test1.txt"}, 140},
		{"part 1 test", args{path: "day12-test2.txt"}, 772},
		{"part 1 test", args{path: "day12-test3.txt"}, 1930},
		{"part 1 answer", args{path: "day12.txt"}, 1467094},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.path); got != tt.want {
				t.Errorf("day12Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day12Part2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 test", args{path: "day12-test1.txt"}, 80},
		{"part 2 test", args{path: "day12-test2.txt"}, 436},
		{"part 2 test", args{path: "day12-test3.txt"}, 1206},
		{"part 2 test", args{path: "day12-test4.txt"}, 236},
		{"part 2 test", args{path: "day12-test5.txt"}, 368},
		{"part 2 answer", args{path: "day12.txt"}, 881182},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.path); got != tt.want {
				t.Errorf("day12Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
