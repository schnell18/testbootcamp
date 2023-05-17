package reverse

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"TC1", "Hello, world", "dlrow ,olleH"},
		{"TC2", " ", " "},
		{"TC3", "!12345", "54321!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reverse(tt.arg)
			assert.Nil(err)
			assert.Equal(tt.want, got)
		})
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		// Use f.Add to provide a seed corpus
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		assert := assert.New(t)
		rev, err1 := Reverse(orig)
		doubleRev, err2 := Reverse(rev)
		if err1 != nil || err2 != nil {
			t.Skip("Ignore invalid utf8 string")
		}

		assert.Equal(orig, doubleRev)
		assert.True(utf8.ValidString(orig))
		assert.True(utf8.ValidString(rev))
	})
}
