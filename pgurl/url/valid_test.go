package url

import (
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	type args struct {
		toTestUrl string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "True url https://www.google.com",
			args: args{
				toTestUrl: "https://www.google.com",
			},
			want: true,
		},
		{
			name: "True url http://www.google.com",
			args: args{
				toTestUrl: "http://www.google.com",
			},
			want: true,
		},
		{
			name: "False url htt://www.google.com",
			args: args{
				toTestUrl: "htt://www.google.com",
			},
			want: false,
		},
		{
			name: "False url http:/www.google.com",
			args: args{
				toTestUrl: "http:/www.google.com",
			},
			want: false,
		},
		{
			name: "Empty url",
			args: args{
				toTestUrl: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUrl(tt.args.toTestUrl); got != tt.want {
				t.Errorf("IsValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
