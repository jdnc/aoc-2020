package main 

import (
 "os"
 "bufio"
 "log"
 "fmt"
 "strconv"
 "strings"
)


func main() {
  // read input
  in, err := os.Open("data/day2.txt")

  if err != nil  {
    log.Fatal(err)
    panic("Unable to open input file")
  }
  defer  in.Close()
  
  scanner := bufio.NewScanner(in)
  ret1 := 0
  ret2 := 0
  for scanner.Scan() {
    // basic tokenization in place
    example := strings.TrimSpace(scanner.Text())
    tokens := strings.Split(example, " ")

    if len(tokens) != 3 {
      panic("Tokenization failed!")
    }

    // get expected frequency for char
    freqs := strings.Split(tokens[0], "-")
    if len(freqs) != 2 {
      panic("Frequency extraction failed!")
    }
    minFreq, err := strconv.Atoi(freqs[0])
    maxFreq, err := strconv.Atoi(freqs[1])
    if err != nil {
      panic("str to int coercion failed!")
    }

    // get actual character
    letter := tokens[1][:len(tokens[1]) - 1]
    
    // get password
    password := tokens[2]

    // check if format is valid
    count := strings.Count(password, letter) 
    if count >= minFreq && count <= maxFreq {
      ret1++;
    } 

    pos1 := string(password[minFreq - 1])
    pos2 := string(password[maxFreq - 1])

    if (pos1 == letter || pos2 == letter) && (pos1 != pos2) {
      ret2++;
    }

  } 

  fmt.Println("[part 1] Number of valid passwords: ", ret1)
  fmt.Println("[part 2] Number of valid passwords: ", ret2)
}
