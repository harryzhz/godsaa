package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testACTrie = NewACTrie()

func TestACTrie_Inserts(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "inserts",
			args: args{
				words: []string{"abc", "bc", "abcd"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testACTrie.Inserts(tt.args.words...)
		})
	}
	testACTrie.trie.Reset()
}

func TestACTrie_Match(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "match1",
			args: args{
				text: "fabcdf",
			},
			want: []string{
				"abc", "bc", "abcd",
			},
		},
		{
			name: "match2",
			args: args{
				text: "你个sb是傻逼逼吧",
			},
			want: []string{
				"sb",
				"傻逼",
				"逼",
			},
		},
	}
	testACTrie.Inserts("abc", "bc", "abcd")
	testACTrie.Inserts("fuck", "shit", "cao", "艹", "操", "sb", "傻逼", "煞笔", "逼")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testACTrie.Match(tt.args.text)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_buildFailPointer(t *testing.T) {
	type args struct {
		trie *Trie
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
