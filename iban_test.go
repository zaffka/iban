package iban_test

import (
	"testing"

	"github.com/zaffka/iban"
)

func TestValid(t *testing.T) {
	type args struct {
		probe string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "num",
			args: args{probe: "036020343320"},
			want: false,
		},
		{
			name: "ok",
			args: args{probe: "RS35265100000071202319"},
			want: true,
		},
		{
			name: "changed_digit",
			args: args{probe: "RS35265100000071202310"},
			want: false,
		},
		{
			name: "wrong_code",
			args: args{probe: "KZ22551B229629855USD"},
			want: false,
		},
		{
			name: "cy",
			args: args{probe: "CY30905000010000001000000469"},
			want: true,
		},
		{
			name: "wrong_num",
			args: args{probe: "CY30905000010000001030000469"},
			want: false,
		},
		{
			name: "short_ok",
			args: args{probe: "RS351"},
			want: true,
		},
		{
			name: "too_short",
			args: args{probe: "RS35"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iban.Valid(tt.args.probe); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
