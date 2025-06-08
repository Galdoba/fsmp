package subtitle

import "testing"

func Test_parseTimestamps(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line    string
		want    float64
		want2   float64
		wantErr bool
	}{
		{
			name:    "correct 1",
			line:    "00:00:59,612 --> 00:01:02,141",
			want:    59.612,
			want2:   62.141,
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, gotErr := parseTimestamps(tt.line)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parseTimestamps() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parseTimestamps() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("parseTimestamps() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("parseTimestamps() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
