package main

import "testing"

func Test_calc(t *testing.T) {
	type args struct {
		strArr [3]string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"default",
			args{strArr: [3]string{"15:01:00", "18:09:45", "15:01:40"}}, "18:10:05"},
		{"next day response",
			args{strArr: [3]string{"23:55:00", "00:20:00", "00:05:00"}}, "00:25:00"},
		{"next day response",
			args{strArr: [3]string{"23:59:00", "00:20:00", "00:01:00"}}, "00:21:00"},
		{"next day response",
			args{strArr: [3]string{"00:00:00", "05:00:00", "00:00:00"}}, "17:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := calc(tt.args.strArr); gotRes != tt.wantRes {
				t.Errorf("calc() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
