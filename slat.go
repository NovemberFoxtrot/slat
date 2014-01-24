package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

type bigram struct {
	left  rune
	right rune
}

type bigramvector []bigram

func (b bigramvector) String() string {
	var result string

	for _, v := range b {
		result += string(v.left)
	}

	return result
}

func (b bigram) String() string {
	return string(b.left) + string(b.right)
}

func jaccard(x, y []bigram) float64 {
	intersection := make(map[bigram]int)
	union := make(map[bigram]int)

	for _, xrune := range x {
		for _, yrune := range y {
			if (xrune.left == yrune.left) && (xrune.right == yrune.right) {
				intersection[xrune] += 1
				union[xrune] += 1
			} else {
				union[xrune] += 1
				union[yrune] += 1
			}
		}
	}

	distance := 1 - ((float64(len(union)) - float64(len(intersection))) / float64(len(union)))

	return distance
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	scanner := bufio.NewScanner(os.Stdin)

	bigramsarray := make([]bigramvector, 0)

	freq := make(map[rune]int)
	total := 0

	for scanner.Scan() {
		bigrams := make(bigramvector, 0)

		someline := scanner.Text()

		var prev rune
		var endrune rune

		for _, r := range someline {
			freq[r] += 1
			total += 1
			bigrams = append(bigrams, bigram{prev, r})
			prev = r
		}

		bigrams = append(bigrams, bigram{prev, endrune})

		bigramsarray = append(bigramsarray, bigrams)
	}

	lexd := 1.0 * float64(len(freq)) / float64(total)
	avg := 1.0 * float64(total) / float64(len(freq))

	fmt.Println(len(freq), freq, total, "lexd", lexd, "avg", avg)

	results := make(map[string]float64)

	for i, x := range bigramsarray {
		for _, y := range bigramsarray[:i] {
			var k1 string
			var k2 string

			if results[k1+k2] > 0.0 || results[k2+k1] > 0.0 {
				continue
			}

			for _, v := range x {
				k1 += string(v.left) + string(v.right)
			}
			for _, v := range y {
				k2 += string(v.left) + string(v.right)
			}

			distance := jaccard(x, y)
			if distance > 0.5 && distance < 1.0 {
				fmt.Println(distance)
				fmt.Println(x)
				fmt.Println(y)
			}
		}
	}
}
