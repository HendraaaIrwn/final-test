package main

func main() {
	printPyramid(7, 1)
	printPyramid(7, 5)

}

func printPyramid(n int, startFrom int) {
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= i; j++ {
			print(startFrom + i + j)
		}
		println()
	}
}
