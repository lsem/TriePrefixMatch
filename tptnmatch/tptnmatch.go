package tptnmatch

import "unicode/utf8"

type TrieEdge struct {
	symbol rune // Keep symbol also in edge in order to not jump in memory much and check symbol without following edges
	node   *TrieNode
}

type TrieNode struct {
	edges         []TrieEdge
	parent        *TrieNode
	rune          rune
	patternEnding bool
}

type Trie struct {
	root TrieNode
}

func (self *TrieNode) EdgeForRune(r rune) *TrieEdge {
	// Iterating with index appeared to be faster than with range.
	for idx := 0; idx < len(self.edges); idx++ {
		if self.edges[idx].symbol == r {
			return &self.edges[idx]
		}
	}
	return nil
}

func (self *TrieNode) IsLeave() bool {
	return len(self.edges) == 0
}

func (self *TrieNode) GetCurrentPattern() string {
	if !self.IsLeave() && !self.patternEnding {
		panic("Unexpected usage")
	}
	sequence := make([]rune, 0, 128)
	for node := self; node != nil; node = node.parent {
		sequence = append(sequence, node.rune)
	}
	// Note, last rune in sequence is root's 0 rune, just not take into consideration it.
	sequence = sequence[:len(sequence)-1 ]

	// Sequence contains reversed pattern, read runes backwards and construct pattern string
	reversedSequence := make([]rune, len(sequence))
	for i := 0; i < len(sequence); i++ {
		reversedSequence[i] = sequence[len(sequence)-1-i]
	}
	return string(reversedSequence)
}

func BuildTrie(patterns []string) Trie {
	result := Trie{}
	root := &result.root
	for _, pattern := range patterns {
		currentNode := root
		for _, symbol := range pattern {
			// Try to find existing edge for symbol
			if symbolEdge := currentNode.EdgeForRune(symbol); symbolEdge != nil {
				currentNode = symbolEdge.node
			} else {
				// not found, insert new node.
				newNode := TrieNode{parent: currentNode, rune: symbol}
				newEdge := TrieEdge{symbol: symbol, node: &newNode}
				currentNode.edges = append(currentNode.edges, newEdge)
				currentNode = &newNode
			}
		}
		currentNode.patternEnding = true
	}
	return result
}

type MatchCallback func(string)

func PrefixTrieMatching(text string, textSize int, trie Trie, matchCb MatchCallback) {
	currentNode := &trie.root
	for runeIdx, r := range text {
		if edge := currentNode.EdgeForRune(r); edge != nil {
			currentNode = edge.node
			// Either leave of tree or matched last character of text with last character of pattern
			if currentNode.IsLeave() || ((runeIdx == textSize-1) && currentNode.patternEnding) {
				matchCb(currentNode.GetCurrentPattern())
				return
			}
		} else if currentNode.patternEnding {
			matchCb(currentNode.GetCurrentPattern())
			return
		} else {
			return
		}
	}
}

func MatchTextAgainstTrie(text string, trie Trie, matchCb MatchCallback) {
	// Pass size to not make our algorithm O(N^2)
	textSize := utf8.RuneCountInString(text)
	for len(text) > 0 {
		PrefixTrieMatching(text, textSize, trie, matchCb)
		// Trim first character of the text
		_, runeSize := utf8.DecodeRuneInString(text)
		text = text[runeSize:]
		textSize--
	}
}
