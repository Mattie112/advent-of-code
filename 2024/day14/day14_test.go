package main

import "testing"

func Test_day14Part1(t *testing.T) {
	type args struct {
		path    string
		width   int
		height  int
		seconds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 1 test", args{path: "day14-test.txt", width: 11, height: 7, seconds: 100}, 12},
		{"part 1 answer", args{path: "day14.txt", width: 101, height: 103, seconds: 100}, 230436441},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.path, tt.args.width, tt.args.height, tt.args.seconds); got != tt.want {
				t.Errorf("day14Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day14Part2(t *testing.T) {
	type args struct {
		path    string
		width   int
		height  int
		seconds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part 2 answer", args{path: "day14.txt", width: 101, height: 103, seconds: 100}, 8270},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.path, tt.args.width, tt.args.height, tt.args.seconds); got != tt.want {
				t.Errorf("day14Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
