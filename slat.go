package main

import (
  "fmt"
  "bufio"
  "runtime"
  "os"
)

type bigram struct {
  left rune
  right rune
}

func (b bigram) String() string {
  return string(b.left) + string(b.right)
}

/*
(
len(X.union(Y)) 
- 
len(X.intersection(Y))
)
/
float(len(X.union(Y)))
*/

func jaccard(x, y []bigram) {
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

  if distance > 0.5 && distance < 1.0 {
    fmt.Println(distance)
    fmt.Println(x)
    fmt.Println(y)
  }
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  scanner := bufio.NewScanner(os.Stdin)

  bigramsarray := make([][]bigram, 0)

  freq := make(map[rune]int)
  total := 0

  for scanner.Scan() {
    bigrams := make([]bigram, 0)

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

  lexd := 1.0*float64(len(freq))/float64(total)
  avg := 1.0*float64(total)/float64(len(freq))

  fmt.Println(len(freq), freq, total, "lexd", lexd, "avg", avg)

  for _, x := range bigramsarray {
    for _, y := range bigramsarray {
      jaccard(x, y)
    }
  }
}
