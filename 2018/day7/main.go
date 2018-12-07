package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
)

const inputPath = "input.txt"

var NodeRE = regexp.MustCompile("Step ([a-zA-Z]) must be finished before step ([a-zA-Z]) can begin")

type Node struct {
	Name     rune
	Children []*Node
	Parents  []*Node
}

type Tree struct {
	Nodes []*Node
}

type NodeSlice []*Node

func (n NodeSlice) Len() int {
	return len(n)
}

func (n NodeSlice) Less(i, j int) bool {
	return n[i].Name < n[j].Name
}

func (n NodeSlice) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *Node) Exists(t *Tree) bool {
	for _, nn := range t.Nodes {
		if nn.Name == n.Name {
			return true
		}
	}

	return false
}

func (n *Node) HasParent(r rune) bool {
	for _, nn := range n.Parents {
		if nn.Name == r {
			return true
		}
	}

	return false
}

func (n *Node) HasChild(r rune) bool {
	for _, nn := range n.Children {
		if nn.Name == r {
			return true
		}
	}

	return false
}

func (t *Tree) HasNodeName(r rune) bool {
	for _, n := range t.Nodes {
		if n.Name == r {
			return true
		}
	}

	return false
}

func (t *Tree) GetNodeName(r rune) *Node {
	for _, n := range t.Nodes {
		if n.Name == r {
			return n
		}
	}

	return nil
}

func NewTree() *Tree {
	t := &Tree{
		Nodes: make([]*Node, 0),
	}

	return t
}

func (t *Tree) Update(n *Node) {
	for _, curNode := range t.Nodes {
		if curNode.Name == n.Name {
			curNode.Children = n.Children
			curNode.Parents = n.Parents
			return
		}
	}

	// Otherwise, no match, we'll just add it to our nodes
	t.Nodes = append(t.Nodes, n)
}

func (t *Tree) Insert(line string) error {
	matches := NodeRE.FindStringSubmatch(line)
	if len(matches) != 3 || len(matches[1]) != 1 || len(matches[2]) != 1 {
		return errors.New("invalid line")
	}

	parent := rune(matches[1][0])
	child := rune(matches[2][0])

	parentNode := t.GetNodeName(parent)
	if parentNode == nil {
		parentNode = &Node{
			Name:     parent,
			Children: make([]*Node, 0),
		}
	}

	childNode := t.GetNodeName(child)
	if childNode == nil {
		childNode = &Node{
			Name:     child,
			Children: make([]*Node, 0),
		}
	}

	// Update child and parents so they refer to each other
	if !parentNode.HasChild(child) {
		parentNode.Children = append(parentNode.Children, childNode)
	}

	if !childNode.HasParent(parent) {
		childNode.Parents = append(childNode.Parents, parentNode)
	}

	t.Update(parentNode)
	t.Update(childNode)

	return nil
}

func solve(t *Tree) (string, error) {
	// Step1: find the node without any parents
	sort.Sort(NodeSlice(t.Nodes))

	var startingNode *Node
	for _, n := range t.Nodes {
		if len(n.Parents) == 0 {
			startingNode = n
			break
		}
	}

	if startingNode == nil {
		return "", errors.New("could not find starting node")
	}

	// Let's us know if we've used the rune yet
	m := make(map[rune]bool)
	m[startingNode.Name] = true

	// Available nodes to pick from
	availableChoices := NodeSlice(startingNode.Children)

	// Keep track of our solution
	var solution string
	solution += string(startingNode.Name)

Beginning:
	// Sort the array by name, so we iterate in order
	sort.Sort(availableChoices)

	for _, n := range availableChoices {
		// if we've seen this node before, we can skip it
		if _, ok := m[n.Name]; ok {
			continue
		}

		// Have to make sure all parents have been seen
		var parentsSeen int
		for _, p := range n.Parents {
			fmt.Printf("%s has parent %s\n", string(n.Name), string(p.Name))
			if _, ok := m[p.Name]; !ok {
				continue
			}
			parentsSeen++
		}

		if parentsSeen != len(n.Parents) {
			continue
		}

		solution += string(n.Name)
		availableChoices = append(availableChoices, n.Children...)
		m[n.Name] = true

		goto Beginning
	}

	return solution, nil
}

func gatherInputs(input io.Reader) (*Tree, error) {
	tree := NewTree()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		err := tree.Insert(line)
		if err != nil {
			return nil, err
		}
	}

	return tree, nil
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tree, err := gatherInputs(file)
	if err != nil {
		panic(err)
	}

	result, err := solve(tree)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
