package common

import (
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		name string
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case-1", args{"ls", []string{"-l", "-ha"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(string(got))
		})
	}
}
