package main

import "fmt"

type node struct {
	isEnd    bool
	children [26]*node
}

type trie struct {
	root *node
}

func (t *trie) insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *trie) search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}

	return false
}

func initTrie() *trie {
	result := &trie{root: &node{}}
	return result

}

func main() {
	trie := initTrie()
	fmt.Println(trie)
	trie.insert("toto")
	fmt.Println(trie.search("toto"))
}
