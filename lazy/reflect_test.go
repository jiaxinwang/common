package lazy

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_foreignOfModel(t *testing.T) {
	type args struct {
		inter interface{}
	}
	tests := []struct {
		name string
		args args
		want [][3]string
	}{
		{"dog", args{inter: Dog{}}, [][3]string{[3]string{`Profile`, `profiles`, `dog_id`}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foreignOfModel(tt.args.inter); !cmp.Equal(got, tt.want) {
				t.Errorf("foreignOfModel() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
