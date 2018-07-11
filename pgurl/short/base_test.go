package short

import "testing"

func Test_getAesKey(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "develop key",
			want: "TesttestTesttest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAesKey(); got != tt.want {
				t.Errorf("getAesKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
