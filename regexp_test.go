package repro

import (
	"fmt"
	re2 "github.com/wasilibs/go-re2/experimental"
	"rsc.io/binaryregexp"
	"testing"
)

func TestUnicode(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    bool
	}{
		{
			pattern: "ハロー",
			input:   "ハローワールド",
			want:    true,
		},
		{
			pattern: "ハロー",
			input:   "グッバイワールド",
			want:    false,
		},
		{
			pattern: `\xac\xed\x00\x05`,
			input:   "\xac\xed\x00\x05t\x00\x04test",
			want:    true,
		},
		{
			pattern: `\xac\xed\x00\x05`,
			input:   "\xac\xed\x00t\x00\x04test",
			want:    false,
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(fmt.Sprintf("%s/%s", tt.pattern, tt.input), func(t *testing.T) {
			t.Run("re2", func(t *testing.T) {
				re := re2.MustCompileLatin1(tt.pattern)
				if re.MatchString(tt.input) != tt.want {
					t.Errorf("MatchString(%q) = %v, want %v", tt.input, !tt.want, tt.want)
				}
			})
			t.Run("binaryregexp", func(t *testing.T) {
				re := binaryregexp.MustCompile(tt.pattern)
				if re.MatchString(tt.input) != tt.want {
					t.Errorf("MatchString(%q) = %v, want %v", tt.input, !tt.want, tt.want)
				}
			})
		})
	}
}
