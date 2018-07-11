package short

import "testing"

func TestShortURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			args:    args{url: "https://www.google.com"},
			want:    "asdbsa",
			wantErr: false,
		},
		{
			name:    "failed",
			args:    args{url: "httkps://www.google.com"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShortURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
