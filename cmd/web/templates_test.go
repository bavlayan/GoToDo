package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	/*tm := time.Date(1995, 05, 02, 22, 0, 0, 0, time.UTC)
	hd := humanDate(tm)

	if hd != "02 May 1995 at 22:00" {
		t.Errorf("want %q; got %q", "02 May 1995 at 22:00", hd)
	}*/

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(1995, 05, 02, 22, 0, 0, 0, time.UTC),
			want: "02 May 1995 at 22:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(1995, 05, 02, 22, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "02 May 1995 at 21:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}
