package main

import "fmt"

// 与えられたn個のペアの括弧から、すべての正しく形成された括弧の組み合わせを生成する関数を実装することを求める。例えば、n=3の場合、すべての可能な組み合わせは以下の5
// "((()))"
// "(()())"
// "(())()"
// "()(())"
// "()()()"

// n=1の場合
// "()"
// n=2の場合
// "(())"
// "()()"
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, current string, open int, close int, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}
	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}
