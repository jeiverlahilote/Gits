package main

import (
	"fmt"
)

// ---------------------- Soal 1: A000124 ----------------------

func a000124(n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, (i*(i+1))/2+1)
	}
	return result
}

func printA000124Examples() {
	fmt.Println("\n--- Soal 1: A000124 ---")
	tests := []int{5, 7, 10}
	for _, t := range tests {
		fmt.Printf("Input: %d -> Output: %v\n", t, a000124(t))
	}
}

// ---------------------- Soal 2: Dense Ranking ----------------------

func removeDuplicates(scores []int) []int {
	unique := []int{}
	seen := make(map[int]bool)
	for _, score := range scores {
		if !seen[score] {
			unique = append(unique, score)
			seen[score] = true
		}
	}
	return unique
}

func denseRanking(leaderboard []int, player []int) []int {
	unique := removeDuplicates(leaderboard)
	result := []int{}
	n := len(unique)

	for _, score := range player {
		left, right := 0, n-1
		for left <= right {
			mid := (left + right) / 2
			if score == unique[mid] {
				left = mid
				break
			} else if score > unique[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		result = append(result, left+1)
	}
	return result
}

func printDenseRankingExamples() {
	fmt.Println("\n--- Soal 2: Dense Ranking ---")
	testCases := []struct {
		leaderboard []int
		player      []int
	}{
		{[]int{100, 100, 50, 40, 40, 20, 10}, []int{5, 25, 50, 120}},
		{[]int{100, 90, 90, 80, 75, 60}, []int{50, 65, 77, 90, 102}},
		{[]int{200, 200, 150, 100, 100, 50}, []int{150, 99, 201}},
	}
	for _, tc := range testCases {
		fmt.Printf("Leaderboard: %v\nPlayer: %v -> Ranking: %v\n", tc.leaderboard, tc.player, denseRanking(tc.leaderboard, tc.player))
	}
}

// ---------------------- Soal 3: Balanced Bracket ----------------------

func isBalanced(s string) string {
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func printBalancedBracketExamples() {
	fmt.Println("\n--- Soal 3: Balanced Bracket ---")
	tests := []string{
		"{[()]}",
		"{{([])}}",
		"{[(])}",
		"((({[()]})[]))",
		"[({(())}[()])]",
		"{[((){})]}",
	}
	for _, t := range tests {
		fmt.Printf("Input: %s -> Output: %s\n", t, isBalanced(t))
	}
}

// ---------------------- Main ----------------------

func main() {
	printA000124Examples()
	printDenseRankingExamples()
	printBalancedBracketExamples()
}
