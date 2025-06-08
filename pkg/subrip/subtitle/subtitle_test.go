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
		{name: "correct simple 1", line: "00:00:59,612 --> 00:01:02,141", want: 59.612, want2: 62.141, wantErr: false},
		{name: "correct simple 2", line: "00:00:59,612   -->       00:01:02,141", want: 59.612, want2: 62.141, wantErr: false},
		{name: "incorrect simple 1", line: "00:00:59.612 --> 00:01:02.141", want: 59.612, want2: 62.141, wantErr: true},
		{name: "incorrect simple 2", line: "00:00:59.612 ---> 00:01:02.141", want: 59.612, want2: 62.141, wantErr: true},
		{name: "correct complex 1", line: "01:07:35,680 --> 01:07:40,840", want: 4055.68, want2: 4060.840, wantErr: false},
		{name: "incorrect complex 1", line: "01:07:35,68 --> 01:07:40,84", want: 4055.68, want2: 4060.840, wantErr: true},
		// TODO: Add test cases.
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

func Test_timestamp(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		seconds float64
		want    string
	}{
		{name: "zero", seconds: 0, want: "00:00:00,000"},
		{name: "0.001", seconds: 0.001, want: "00:00:00,001"},
		{name: "0.0001", seconds: 0.0001, want: "00:00:00,000"},
		{name: "0.031", seconds: 0.031, want: "00:00:00,031"},
		{name: "0.301", seconds: 0.301, want: "00:00:00,301"},
		{name: "0.310", seconds: 0.310, want: "00:00:00,310"},
		{name: "0.31", seconds: 0.31, want: "00:00:00,310"},
		{name: "42.31", seconds: 42.31, want: "00:00:42,310"},
		{name: "162.31", seconds: 162.31, want: "00:02:42,310"},
		{name: "762.31", seconds: 762.31, want: "00:12:42,310"},
		{name: "3605.1", seconds: 3605.1, want: "01:00:05,100"},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := timestamp(tt.seconds)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
