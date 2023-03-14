package main

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"", args{text: "(({[]}))"}, "yes"},
		{"", args{text: "((((()))))"}, "yes"},
		{"", args{text: "("}, "no"},
		{"", args{text: ")"}, "no"},
		{"", args{text: "({)"}, "no"},
		{"", args{text: "({)}"}, "no"},
		{"", args{text: "([][)]"}, "no"},
		{"", args{text: "{{{}}"}, "no"},
		{"", args{text: "{{{}}}"}, "yes"},
		{"", args{text: ")("}, "no"},
		{"", args{text: "(()))"}, "no"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := calculate(tt.args.text); gotRes != tt.wantRes {
				t.Errorf("calculate() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
