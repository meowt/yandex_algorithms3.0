package main

import (
	"reflect"
	"testing"
)

func Test_parseData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		args      args
		wantGraph [][]int
	}{
		{
			args:      args{data: []byte("4 5\n2 2\n3 4\n2 3\n1 3\n2 4\n\n")},
			wantGraph: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGraph := parseData(tt.args.data); !reflect.DeepEqual(gotGraph, tt.wantGraph) {
				t.Errorf("parseData() = %v, want %v", gotGraph, tt.wantGraph)
			}
		})
	}
}
