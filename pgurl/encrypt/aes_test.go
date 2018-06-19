package encrypt

import (
	"reflect"
	"testing"
)

func Test_addBase64Padding(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "len %4 == 1",
			args: args{
				value: "a",
			},
			want: "a===",
		},
		{
			name: "len %4 == 2",
			args: args{
				value: "aa",
			},
			want: "aa==",
		},
		{
			name: "len %4 == 3",
			args: args{
				value: "aaa",
			},
			want: "aaa=",
		},
		{
			name: "len %4 == 0",
			args: args{
				value: "aaaa",
			},
			want: "aaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBase64Padding(tt.args.value); got != tt.want {
				t.Errorf("addBase64Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeBase64Padding(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "len %4 == 1",
			args: args{
				value: "a===",
			},
			want: "a",
		},
		{
			name: "len %4 == 2",
			args: args{
				value: "aa==",
			},
			want: "aa",
		},
		{
			name: "len %4 == 3",
			args: args{
				value: "aaa=",
			},
			want: "aaa",
		},
		{
			name: "len %4 == 0",
			args: args{
				value: "aaaa",
			},
			want: "aaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBase64Padding(tt.args.value); got != tt.want {
				t.Errorf("removeBase64Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPad(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test 1",
			args: args{
				src: []byte("1234"),
			},
			want: []byte{49, 50, 51, 52, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		},
		{
			name: "test 2",
			args: args{
				src: []byte("abcde"),
			},
			want: []byte{97, 98, 99, 100, 101, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		},
		{
			name: "test 3",
			args: args{
				src: []byte("一二三四五六七"),
			},
			want: []byte{228, 184, 128, 228, 186, 140, 228, 184, 137, 229, 155, 155, 228, 186, 148, 229, 133, 173, 228, 184, 131, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pad(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnpad(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				src: []byte{49, 50, 51, 52, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
			},
			want: []byte("1234"),
		},
		{
			name: "test 2",
			args: args{
				src: []byte{97, 98, 99, 100, 101, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
			},
			want: []byte("abcde"),
		},
		{
			name: "test 3",
			args: args{
				src: []byte{228, 184, 128, 228, 186, 140, 228, 184, 137, 229, 155, 155, 228, 186, 148, 229, 133, 173, 228, 184, 131, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
			},
			want: []byte("一二三四五六七"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpad(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unpad() = %v, want %v", got, tt.want)
			}
		})
	}
}
