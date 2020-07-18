package lazy

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_separateParams(t *testing.T) {
	type args struct {
		whole Params
		keys  []string
	}
	tests := []struct {
		name          string
		args          args
		wantSeparated Params
		wantRemain    Params
	}{
		{
			"case-1",
			args{whole: Params{"a": []string{`v1`, `v2`}}, keys: []string{`a`}},
			Params{"a": []string{`v1`, `v2`}},
			Params{},
		},
		{
			"case-2",
			args{whole: Params{"a": []string{`v1`, `v2`}}, keys: []string{}},
			Params{},
			Params{"a": []string{`v1`, `v2`}},
		},
		{
			"case-3",
			args{whole: Params{"a": []string{`v1`, `v2`}, "b": []string{`v3`, `v4`}}, keys: []string{`a`}},
			Params{"a": []string{`v1`, `v2`}},
			Params{"b": []string{`v3`, `v4`}},
		},
		{
			"case-4",
			args{whole: Params{"a": []string{`v1`, `v2`}, "b": []string{`v3`, `v4`}}, keys: []string{`c`}},
			Params{},
			Params{"a": []string{`v1`, `v2`}, "b": []string{`v3`, `v4`}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSeparated, gotRemain := separateParams(tt.args.whole, tt.args.keys...)
			if !cmp.Equal(gotSeparated, tt.wantSeparated) {
				t.Errorf("separateParams() gotSeparated = %v, want %v\ndiff=%v", gotSeparated, tt.wantSeparated, cmp.Diff(gotSeparated, tt.wantSeparated))
			}
			if !cmp.Equal(gotRemain, tt.wantRemain) {
				t.Errorf("separateParams() gotRemain = %v, want %v\ndiff=%v", gotRemain, tt.wantRemain, cmp.Diff(gotRemain, tt.wantRemain))
			}
		})
	}
}

func Test_separatePrefixParams(t *testing.T) {
	type args struct {
		whole  Params
		prefix string
	}
	tests := []struct {
		name          string
		args          args
		wantSeparated Params
		wantRemain    Params
	}{
		{
			"case-1",
			args{whole: Params{"p_a": []string{`v1`, `v2`}}, prefix: `p_`},
			Params{"p_a": []string{`v1`, `v2`}},
			Params{},
		},
		{
			"case-2",
			args{whole: Params{"a": []string{`v1`, `v2`}}, prefix: `p_`},
			Params{},
			Params{"a": []string{`v1`, `v2`}},
		},
		{
			"case-3",
			args{whole: Params{"p_a": []string{`v1`, `v2`}, "b": []string{`v3`, `v4`}}, prefix: `p_`},
			Params{"p_a": []string{`v1`, `v2`}},
			Params{"b": []string{`v3`, `v4`}},
		},
		{
			"case-4",
			args{whole: Params{"p_a": []string{`v1`, `v2`}, "p_b": []string{`v3`, `v4`}}, prefix: `p_`},
			Params{"p_a": []string{`v1`, `v2`}, "p_b": []string{`v3`, `v4`}},
			Params{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSeparated, gotRemain := separatePrefixParams(tt.args.whole, tt.args.prefix)
			if !cmp.Equal(gotSeparated, tt.wantSeparated) {
				t.Errorf("separatePrefixParams() gotSeparated = %v, want %v\ndiff=%v", gotSeparated, tt.wantSeparated, cmp.Diff(gotSeparated, tt.wantSeparated))
			}
			if !cmp.Equal(gotRemain, tt.wantRemain) {
				t.Errorf("separatePrefixParams() gotRemain = %v, want %v\ndiff=%v", gotRemain, tt.wantRemain, cmp.Diff(gotRemain, tt.wantRemain))
			}
		})
	}
}
