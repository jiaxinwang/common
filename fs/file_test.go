package fs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFileExist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExist(tt.args.name); got != tt.want {
				t.Errorf("FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadlines(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name      string
		args      args
		wantLines []string
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLines, err := Readlines(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Readlines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotLines, tt.wantLines) {
				t.Errorf("Readlines() = %v, want %v\ndiff=%v", gotLines, tt.wantLines, cmp.Diff(gotLines, tt.wantLines))
			}
		})
	}
}

func TestSafeCopy(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case-0", args{"./file_test.go", "./file_test_copy.bak"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SafeCopy(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("SafeCopy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
