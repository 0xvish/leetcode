package main

func main() {
	answer := maxScore("011101")
	println(answer)
}

func maxScore(s string) int {
	score := 0
	n := len(s)
	for i := 1; i < n; i++ {
		zeros, _ := countChar(s[0:i])
		_, ones := countChar(s[i:])
		sum := zeros + ones
		if score < sum {
			score = sum
		}
	}

	return score
}

func countChar(s string) (int, int) {
	zeros, ones := 0, 0

	for _, char := range s {
		if char == '0' {
			zeros++
		}
		if char == '1' {
			ones++
		}

	}
	return zeros, ones

}