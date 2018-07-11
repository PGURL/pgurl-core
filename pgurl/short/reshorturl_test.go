package short

import "testing"

func TestReShort(t *testing.T) {
	baseURL := "https://www.google.com"
	shortURL, _ := ShortURL(baseURL)
	type args struct {
		urlKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "exist url",
			args: args{
				urlKey: shortURL,
			},
			want:    baseURL,
			wantErr: false,
		},
		{
			name: "not exist url",
			args: args{
				urlKey: "https://hello.com",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReShort(tt.args.urlKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReShort() = %v, want %v", got, tt.want)
			}
		})
	}
}
