package main

import "testing"

func Test_calc(t *testing.T) {
	type args struct {
		slc []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// TODO: Add test cases.
		{"1", args{slc: []int{1, 1}}, 1},
		{"2", args{slc: []int{1, 2, 3, 4}}, 2},
		{"3", args{slc: []int{1, 2, 2, 4}}, 1},
		{"4", args{slc: []int{1, 1, 2, 2}}, 2},
		{"5", args{slc: []int{1, 1, 2, 2, 3, 3}}, 3},
		{"6", args{slc: []int{1, 2, 2, 2, 3, 3}}, 2},
		{"7", args{slc: []int{1, 10, 2, 6, 3, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := calculate(tt.args.slc); gotRes != tt.wantRes {
				t.Errorf("calculate() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
