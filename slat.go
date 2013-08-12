package main

import (
  "fmt"
  "bufio"
  "runtime"
  "os"
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  scanner := bufio.NewScanner(os.Stdin)

  freq := make(map[rune]int)
  total := 0

  for scanner.Scan() {
    someline := scanner.Text()

    for i, r := range someline {
      freq[r] += 1
      total += 1

      fmt.Println(i, len(someline))
    }
  }

  lexd := 1.0*float64(len(freq))/float64(total)
  avg := 1.0*float64(total)/float64(len(freq))

  fmt.Println(len(freq), freq, total, "lexd", lexd, "avg", avg)
}
