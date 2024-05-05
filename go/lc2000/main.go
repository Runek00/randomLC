package main

func reversePrefix(word string, ch byte) string {
	w := []byte(word)
	j := 0

	for k, c := range w {
		if c == ch {
			j = k
			break
		}
	}
	for i := 0; i <= j/2; i++ {
		t := w[i]
		w[i] = w[j-i]
		w[j-i] = t
	}
	return string(w)
}

func main() {}
