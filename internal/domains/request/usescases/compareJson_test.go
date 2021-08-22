package usescases

import (
	"io"
	"strings"
	"testing"
)

func TestCompareJSON_Compare(t *testing.T) {
	type args struct {
		a io.Reader
		b io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "equal json",
			args: args{
				a: strings.NewReader(`{"x": ["y",42]}`),
				b: strings.NewReader(`{"x": ["y",42]}`),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "not equal json",
			args: args{
				a: strings.NewReader(`{"x": ["y",42]}`),
				b: strings.NewReader(`{"a": ["y",42]}`),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "json compare \"\"",
			args: args{
				a: strings.NewReader(`{"x": ["y",42]}`),
				b: strings.NewReader(``),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "\"\" compare \"\"",
			args: args{
				a: strings.NewReader(``),
				b: strings.NewReader(``),
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompareJSON{}
			got, err := c.Compare(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
