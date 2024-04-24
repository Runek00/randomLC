package main

func nbonacci(N int) (func(int) int){
	cache := []int{0}
	for i := 1; i < N; i++ {
		cache = append(cache, 1)
	}

	var Out func(int) int
	Out = func (n int) int {
		if n < len(cache) {
			return cache[n]
		}
		output := 0
		for i := 1; i <= N; i++ {
			output += Out(n - i)
		}
		cache = append(cache, output)
		return output
	}

	return Out

}

func tribonacci(n int) int {
	return nbonacci(3)(n)
}

func main() {
	println(tribonacci(4))
	println(tribonacci(25))
}
