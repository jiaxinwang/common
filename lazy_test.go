package common

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestLazy(t *testing.T) {
	type args struct {
		params map[string]interface{}
	}
	pt, _ := time.Parse("2006-01-02", "2000-01-01")
	tests := []struct {
		name     string
		args     args
		wantEq   map[string][]interface{}
		wantGt   map[string]interface{}
		wantLt   map[string]interface{}
		wantGteq map[string]interface{}
		wantLteq map[string]interface{}
	}{
		{
			"empty",
			args{
				map[string]interface{}{
					"name":           "tom",
					"created_at_lte": pt,
					"w_lt":           0.01,
					"age_gt":         18,
					"p_gte":          32,
					"size":           12,
					"page":           2,
					"offset":         100,
				},
			},
			map[string][]interface{}{"name": []interface{}{"tom"}},
			map[string]interface{}{"age": 18},
			map[string]interface{}{"w": 0.01},
			map[string]interface{}{"p": 32},
			map[string]interface{}{"created_at": pt},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEq, gotGt, gotLt, gotGteq, gotLteq := Lazy(tt.args.params)
			if !cmp.Equal(gotEq, tt.wantEq) {
				t.Errorf("Lazy() gotEq = %v, want %v\ndiff=%v", gotEq, tt.wantEq, cmp.Diff(gotEq, tt.wantEq))
			}
			if !cmp.Equal(gotGt, tt.wantGt) {
				t.Errorf("Lazy() gotGt = %v, want %v\ndiff=%v", gotGt, tt.wantGt, cmp.Diff(gotGt, tt.wantGt))
			}
			if !cmp.Equal(gotLt, tt.wantLt) {
				t.Errorf("Lazy() gotLt = %v, want %v\ndiff=%v", gotLt, tt.wantLt, cmp.Diff(gotLt, tt.wantLt))
			}
			if !cmp.Equal(gotGteq, tt.wantGteq) {
				t.Errorf("Lazy() gotGteq = %v, want %v\ndiff=%v", gotGteq, tt.wantGteq, cmp.Diff(gotGteq, tt.wantGteq))
			}
			if !cmp.Equal(gotLteq, tt.wantLteq) {
				t.Errorf("Lazy() gotLteq = %v, want %v\ndiff=%v", gotLteq, tt.wantLteq, cmp.Diff(gotLteq, tt.wantLteq))
			}
		})
	}
}

func TestLazyTag(t *testing.T) {
	testStruct := struct {
		Name string `lazy:"name"`
	}{}
	type args struct {
		v interface{}
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			"simple",
			args{&testStruct, map[string]interface{}{"Name": "tom"}},
			map[string]interface{}{"name": "tom"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LazyTag(tt.args.v, tt.args.m); !cmp.Equal(got, tt.want) {
				t.Errorf("LazyTag() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
