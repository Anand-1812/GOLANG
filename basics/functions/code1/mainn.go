package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSortWithCycle(m map[string][]string) ([]string, error) {
	var order []string

	// 0 = unvisited, 1 = visiting, 2 = visited
	state := make(map[string]int)

	var visit func(string) error

	visit = func(node string) error {
		switch state[node] {
		case 1:
			return fmt.Errorf("cycle detected at %s", node)
		case 2:
			return nil
		}

		state[node] = 1

		for _, nei := range m[node] {
			if err := visit(nei); err != nil {
				return err
			}
		}

		state[node] = 2
		order = append(order, node)
		return nil
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if state[k] == 0 {
			if err := visit(k); err != nil {
				return nil, err
			}
		}
	}

	return order, nil
}

func main() {
	order, err := topoSortWithCycle(prereqs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
