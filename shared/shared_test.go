package shared

import (
	"reflect"
	"testing"
)

func TestAnalyzeData(t *testing.T) {
	data1 := []string{"satu", "dua", "tiga"}
	data2 := []string{"dua", "tiga", "empat"}
	type args struct {
		array [][]string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				[][]string{
					data1,
					data2,
				},
			},
			want:  []string{"satu", "empat"},
			want1: []string{"dua", "tiga"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AnalyzeData(tt.args.array...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyzeData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AnalyzeData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
