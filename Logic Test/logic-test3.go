package main

func main() {
	inputWord1 := "xhixhix"
	inputWord2 := "mic"
	inputWord3 := "haha"
	inputWord4 := "xxxxyxz"

	println(findWordCounter(inputWord1, "x"))
	println(findWordCounter(inputWord1, "hi"))
	println(findWordCounter(inputWord2, "mic"))
	println(findWordCounter(inputWord3, "ho"))
	println(findWordCounter(inputWord4, "xx"))

}

func findWordCounter(s string, t string) int {
	counter := 0
	for i := 0; i <= len(s)-len(t); i++ {
		if s[i:i+len(t)] == t {
			counter++
		}
	}
	return counter
}
