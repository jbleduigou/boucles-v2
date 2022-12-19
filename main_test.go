package main

import "testing"

func TestPlusOneMinute(t *testing.T) {
	tests := []struct {
		name     string
		timecode string
		want     string
	}{
		{
			want:     "10:29:14:20",
			timecode: "10:28:14:20",
			name:     "Nominal case",
		},
		{
			want:     "10:03:14:20",
			timecode: "10:02:14:20",
			name:     "Padding minutes on two digits",
		},
		{
			want:     "11:00:14:20",
			timecode: "10:59:14:20",
			name:     "Increment hours",
		},
		{
			want:     "06:00:14:20",
			timecode: "05:59:14:20",
			name:     "Increment hours and padding",
		},
		{
			want:     "invalid",
			timecode: "invalid",
			name:     "Garbage in",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOneMinute(tt.timecode); got != tt.want {
				t.Errorf("plusOneMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}
