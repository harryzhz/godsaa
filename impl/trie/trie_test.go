package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testTrie = NewTrie()

func TestTrie_Insert(t1 *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insert1",
			args: args{
				word: "hello",
			},
		},
		{
			name: "insert2",
			args: args{
				word: "head",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			testTrie.Insert(tt.args.word)
		})
	}
	testTrie.Reset()
}

func TestTrie_Inserts(t1 *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insert1",
			args: args{
				words: []string{"hello", "head", "world"},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			testTrie.Inserts(tt.args.words...)
		})
	}
	testTrie.Reset()
}

func TestTrie_Match(t1 *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "match1",
			args: args{
				word: "hello",
			},
			want: true,
		},
	}
	testTrie.Inserts("hello", "head", "world")
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got := testTrie.Match(tt.args.word)
			assert.Equal(t1, tt.want, got)
		})
	}
	testTrie.Reset()
}

func TestTrie_MatchPrefix(t1 *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "match1",
			args: args{
				word: "hello",
			},
			want: true,
		},
		{
			name: "match2",
			args: args{
				word: "hell",
			},
			want: true,
		},
		{
			name: "match3",
			args: args{
				word: "你好",
			},
			want: true,
		},
	}
	testTrie.Inserts("hello", "head", "world", "你好", "你好吗")
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got := testTrie.MatchPrefix(tt.args.word)
			assert.Equal(t1, tt.want, got)
		})
	}
	testTrie.Reset()
}

func TestTrie_search(t1 *testing.T) {
	type args struct {
		word      string
		checkTail bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "search1",
			args: args{
				word:      "hello",
				checkTail: true,
			},
			want: true,
		},
		{
			name: "search2",
			args: args{
				word:      "hell",
				checkTail: true,
			},
			want: false,
		},
		{
			name: "search3",
			args: args{
				word:      "hess",
				checkTail: false,
			},
			want: false,
		},
		{
			name: "search4",
			args: args{
				word:      "hell",
				checkTail: false,
			},
			want: true,
		},
	}
	testTrie.Inserts("hello", "head", "world")
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got := testTrie.search(tt.args.word, tt.args.checkTail)
			assert.Equal(t1, tt.want, got)
		})
	}
	testTrie.Reset()
}
