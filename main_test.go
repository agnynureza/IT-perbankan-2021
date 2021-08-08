package main

import (
	"testing"
)

func Test_grade(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "grade C",
			args: args{
				score: 40,
			},
			want: "C",
		},
		{
			name: "grade B",
			args: args{
				score: 60,
			},
			want: "B",
		},
		{
			name: "grade A",
			args: args{
				score: 80,
			},
			want: "A",
		},
		{
			name: "grade undefine",
			args: args{
				score: 110,
			},
			want: gradeUndefine,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grade(tt.args.score); got != tt.want {
				t.Errorf("grade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "main grade",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
