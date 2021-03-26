package main

// 多叉树节点
type TrieNode struct {
	// 字节点列表
	links []*TrieNode
	R     int
	// 当前单词是否是单词结尾
	isEnd bool
}

func NewTrieNode() *TrieNode {
	t := TrieNode{}
	t.R = 26
	t.links = make([]*TrieNode, t.R)
	return &t
}

func (t TrieNode) ContainsKey(ch byte) bool {
	return t.links[ch-'a'] != nil
}

func (t TrieNode) Get(ch byte) *TrieNode {
	return t.links[ch-'a']
}

func (t *TrieNode) Put(ch byte, node *TrieNode) {
	t.links[ch-'a'] = node
}

func (t *TrieNode) SetEnd() {
	t.isEnd = true
}

func (t *TrieNode) IsEnd() bool {
	return t.isEnd
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		NewTrieNode(),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		// 当前node不存在c这个子节点
		if !p.ContainsKey(c) {
			p.Put(c, NewTrieNode())
		}
		p = p.Get(c)
	}
	p.SetEnd()
}

func (this *Trie) searchPrefix(word string) *TrieNode {
	p := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if p.ContainsKey(c) {
			p = p.Get(c)
		} else {
			return nil
		}
	}
	return p
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.IsEnd()
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.searchPrefix(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
