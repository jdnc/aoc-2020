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
  in, err := os.Open("data/day1.txt")

  if err != nil  {
    log.Fatal(err)
  }

  defer  in.Close()

  var arr[] int

  scanner := bufio.NewScanner(in);
  for scanner.Scan() {
    num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
    if err !=  nil {
      log.Fatal(err)
    }
    arr = append(arr, num)
  }

  // debug
  // fmt.Printf("%v", arr)
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  // process part 1
  for idx, val := range arr {
    targetSum := 2020 - val;
    for i := idx + 1; i < len(arr); i++ {
      if arr[i] == targetSum {
        fmt.Println("==Answer 1==");
        fmt.Println(arr[idx] * arr[i]);
      }
    }
  }
  
  // process part 2
  for i := 0; i < len(arr); i++ {
    for j := i + 1; j < len(arr); j++ {
      for k := j + 1; k < len(arr); k++ {
        if arr[i] + arr[j] + arr[k] == 2020 {
          fmt.Println("==Answer 2==");
          fmt.Println(arr[i] * arr[j] * arr[k]);
        }
      }
    }
  }
}
