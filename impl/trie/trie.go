package trie

type TrieNode struct {
	Data     rune
	IsTail   bool
	Len      int
	Children map[rune]*TrieNode
	Fail     *TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			Data:     '/',
			Children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	rs := []rune(word)
	for _, c := range rs {
		if node.Children == nil {
			node.Children = make(map[rune]*TrieNode)
		}
		if _, ok := node.Children[c]; !ok {
			node.Children[c] = &TrieNode{
				Data: c,
			}
		}
		node = node.Children[c]
	}
	node.IsTail = true
	node.Len = len(rs)
}

func (t *Trie) Inserts(words ...string) {
	for _, w := range words {
		t.Insert(w)
	}
}

// remove the word from the trie
func (t *Trie) Remove(word string) {
	node := t.root
	rs := []rune(word)
	for _, c := range rs {
		if node.Children == nil {
			return
		}
		if _, ok := node.Children[c]; !ok {
			return
		}
		node = node.Children[c]
	}
	if node.IsTail {
		node.IsTail = false
		node.Len = 0
	}
}

// display trie
func (t *Trie) Display() string {
	return ""
}

func (t *Trie) search(word string, checkTail bool) bool {
	node := t.root
	for _, c := range word {
		if node.Children == nil {
			return false
		}
		if _, ok := node.Children[c]; !ok {
			return false
		}
		node = node.Children[c]
	}
	if checkTail {
		return node.IsTail
	}
	return true
}

func (t *Trie) Match(word string) bool {
	return t.search(word, true)
}

func (t *Trie) MatchPrefix(word string) bool {
	return t.search(word, false)
}

func (t *Trie) Reset() {
	t = NewTrie()
}
