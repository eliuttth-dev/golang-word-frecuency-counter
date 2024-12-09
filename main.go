package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main(){
  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Println("Enter a sentence (or type 'exit' to quit):")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    if strings.ToLower(input) == "exit" {
      fmt.Println("Goodbye!")
      break;
    }

    wordCounts := countWordFrecuency(input)
    fmt.Println("Word frecuencies:")

    for word, count := range wordCounts {
      fmt.Printf("%s: %d\n", word, count)
    }

    fmt.Println()
  }
}

func countWordFrecuency(sentence string) map[string]int {
  wordCounts := make(map[string]int)
  words := strings.Fields(sentence)
  
  for _, word := range words {
    word = strings.ToLower(word)
    wordCounts[word]++
  }
  return wordCounts
}
