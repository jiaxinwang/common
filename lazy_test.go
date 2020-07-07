package common

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
}

func TestLazy(t *testing.T) {
	type args struct {
		params map[string][]string
	}
	tests := []struct {
		name     string
		args     args
		wantEq   map[string][]string
		wantGt   map[string]string
		wantLt   map[string]string
		wantGteq map[string]string
		wantLteq map[string]string
	}{
		{
			"empty",
			args{
				map[string][]string{
					"name":           []string{"tom"},
					"created_at_lte": []string{"2000-01-01"},
					"w_lt":           []string{"0.01"},
					"age_gt":         []string{"18"},
					"p_gte":          []string{"32"},
					"size":           []string{"12"},
					"page":           []string{"2"},
					"offset":         []string{"100"},
				},
			},
			map[string][]string{"name": []string{"tom"}},
			map[string]string{"age": "18"},
			map[string]string{"w": "0.01"},
			map[string]string{"p": "32"},
			map[string]string{"created_at": "2000-01-01"},
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

func TestLazyStructMap(t *testing.T) {
	type CustomTime struct {
		Start time.Time `lazy:"Start"`
	}

	newT, _ := time.Parse("20060102", "20200101")
	ct := CustomTime{Start: newT}

	ret := map[string]interface{}{"Start": newT.Format("2006-01-02 15:04:05.000")}

	type args struct {
		src        interface{}
		timeLayout string
	}
	tests := []struct {
		name    string
		args    args
		wantRet map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"time", args{src: ct, timeLayout: "2006-01-02 15:04:05.000"}, ret, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := LazyStructMap(tt.args.src, tt.args.timeLayout)
			if (err != nil) != tt.wantErr {
				t.Errorf("LazyStructMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotRet, tt.wantRet) {
				t.Errorf("LazyStructMap() = %v, want %v\ndiff=%v", gotRet, tt.wantRet, cmp.Diff(gotRet, tt.wantRet))
			}
		})
	}
}
