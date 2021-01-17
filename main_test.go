package bintree

import "testing"

func TestReturnsCorrectResultsFromTrees(t *testing.T) {
	inputs := [][]int{
		{100, 200, 300, 55, 10},
		{100, 200, 300, 55, 100},
	}

	expectedResults := []bool{false, true}

	for idx, tc := range inputs {
		node := Node{}
		for i := 0; i <= len(tc)-2; i++ {
			node.Insert(tc[i])
		}

		searchDigit := tc[len(tc)-1]

		actual := node.Search(searchDigit)
		expected := expectedResults[idx]

		if expected != actual {
			t.Errorf("Failed test case #%d: %t is not %t", idx+1, expected, actual)
		}
	}
}
