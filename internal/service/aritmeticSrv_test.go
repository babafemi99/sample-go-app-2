package service

import "testing"

func Test_arthSrv_Compute(t *testing.T) {
	type args struct {
		operation string
		x         int
		y         int
	}
	tests := []struct {
		name       string
		args       args
		wantOps    string
		wantResult int
	}{
		{
			name: "correct values",
			args: args{
				operation: "addition",
				x:         -10,
				y:         12,
			},
			wantOps:    "addition",
			wantResult: 2,
		},
		{
			name: "correct values",
			args: args{
				operation: "Can you please minus the following numbers together - 13 and 25.",
				x:         -10,
				y:         12,
			},
			wantOps:    "subtraction",
			wantResult: -22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arthSrv{}
			gotOps, gotResult := a.Compute(tt.args.operation, tt.args.x, tt.args.y)
			if gotOps != tt.wantOps {
				t.Errorf("Compute() gotOps = %v, want %v", gotOps, tt.wantOps)
			}
			if gotResult != tt.wantResult {
				t.Errorf("Compute() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
