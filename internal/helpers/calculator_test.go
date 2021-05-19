package helpers

import "testing"

func TestPercentageOneHalf(t *testing.T) {
	type args struct {
		sum float64
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
		wantErr    bool
	}{
		{
			name: "case-1",
			args: args{
				sum: 100.0,
			},
			wantResult: 101.5,
			wantErr:    false,
		},
		{
			name: "case-2",
			args: args{
				sum: 500.0,
			},
			wantResult: 507.5,
			wantErr:    false,
		},
		{
			name: "case-3",
			args: args{
				sum: 0.0,
			},
			wantResult: 0.0,
			wantErr:    true,
		},
		{
			name: "case-4",
			args: args{
				sum: -500.0,
			},
			wantResult: 0.0,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := PercentageOneHalf(tt.args.sum)
			if (err != nil) != tt.wantErr {
				t.Errorf("PercentageOneHalf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("PercentageOneHalf() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
