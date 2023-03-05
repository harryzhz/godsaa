package trie

type ACTrie struct {
	trie *Trie
}

func NewACTrie() *ACTrie {
	return &ACTrie{
		trie: NewTrie(),
	}
}

// buildFailPointer 在 Trie 树上构建失败指针，用于匹配失败时进行跳转
// 失败指针只会出现在上一层，构建过程按次序遍历依次进行遍历, 下一层设置时可以使用上层已设置的失败指针
func buildFailPointer(trie *Trie) {
	if trie == nil {
		return
	}
	// 根节点的失败指针为 nil
	trie.root.Fail = nil
	que := []*TrieNode{}
	que = append(que, trie.root)
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		for c, pc := range p.Children {
			if pc == nil {
				continue
			}
			que = append(que, pc)
			// 根节点的子节点的失败指针指向根节点
			if p == trie.root {
				pc.Fail = p
				continue
			}
			// q 为父节点 p 的失败指针
			// 如果 p 的子节点 pc 对应的字符也可以在 q 的子节点 qc 中找到, 那么 pc 的失败指针就指向 qc
			// 如果找不到 q 跳转到 q 的失败指针节点重复上述操作
			q := p.Fail
			for q != nil {
				qc := q.Children[c]
				if qc != nil {
					pc.Fail = qc
					break
				}
				q = q.Fail
			}
			// 如果最终都没有找到则让 pc 的失败指针指向根节点
			if q == nil {
				pc.Fail = trie.root
			}
		}
	}
}

// Inserts 插入模式单词
func (a *ACTrie) Inserts(words ...string) {
	a.trie.Inserts(words...)
	buildFailPointer(a.trie)
}

// Match 匹配文本
func (a *ACTrie) Match(text string) []string {
	root := a.trie.root
	p := root
	res := make([]string, 0)
	rs := []rune(text)
	for i, c := range rs {
		for p.Children[c] == nil && p != root {
			p = p.Fail
		}
		p = p.Children[c]
		if p == nil {
			p = root
			continue
		}
		tmp := p
		for tmp != root {
			if tmp.IsTail {
				res = append(res, string(rs[i-tmp.Len+1:i+1]))
			}
			tmp = tmp.Fail
		}
	}
	return res
}
