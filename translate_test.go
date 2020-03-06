package common

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var goodbyeMoonmenEN = []string{
	"The worlds can be one together",
	"Cosmos without hatred",
	"Stars like diamonds in your eyes",
	"The ground can be space, space, space, space, space",
	"With feet marching towards a peaceful sky",
	"All the moonmen want things their way",
	"But we make sure they see the sun",
	"Goodbye moonmen",
	"Yeah we say goodbye moonmen",
	"Goodbye moonmen",
	"Goodbye moonmen",
	"Oh goodbye",
	"Cosmos without hatred",
	"Diamond stars of cosmic light",
	"Quasars shine through endless night",
	"And everything is one in the beauty",
	"And now we say goodbye moonmen",
	"Yeah we say goodbye moonmen",
	"Goodbye moonmen",
	"Goodbye moonmen",
	"Oh goodbye",
	"Shut the fuck up about Moonmen!",
}

var goodbyeMoonmenCN = []string{
	"世界可以合而为一",
	"没有仇恨的宇宙",
	"星星就像你眼中的钻石",
	"地面可以是空间，空间，空间，空间，空间",
	"脚踏向宁静的天空",
	"所有的月亮人都希望自己的方式",
	"但是我们确保他们看到了阳光",
	"再见月亮人",
	"是的，我们说再见了",
	"再见月亮人",
	"再见月亮人",
	"哦再见",
	"没有仇恨的宇宙",
	"宇宙光的钻石星",
	"类星体在无尽的夜晚中闪耀",
	"一切都在美丽中合而为一",
	"现在我们说再见了",
	"是的，我们说再见了",
	"再见月亮人",
	"再见月亮人",
	"哦再见",
	"闭嘴他妈的关于月满！",
}

func TestTranslateMultiLines(t *testing.T) {
	type args struct {
		lines []string
		from  string
		to    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"simple sentence",
			args{[]string{"Goodbye moonmen"}, "en", "zh"},
			[]string{"再见月亮人"},
		},
		{
			"lyric slice",
			args{goodbyeMoonmenEN, "en", "zh"},
			goodbyeMoonmenCN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateMultiLines(tt.args.lines, tt.args.from, tt.args.to); !cmp.Equal(got, tt.want) {
				t.Errorf("TranslateMultiLines() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTranslate(t *testing.T) {
	type args struct {
		line string
		from string
		to   string
	}
	tests := []struct {
		name           string
		args           args
		wantTranslated string
		wantErr        bool
	}{
		{"1 word", args{"space", "en", "zh"}, "空间", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTranslated, err := Translate(tt.args.line, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTranslated != tt.wantTranslated {
				t.Errorf("Translate() = %v, want %v", gotTranslated, tt.wantTranslated)
			}
		})
	}
}
