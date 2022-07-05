package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func part2() {
  file, _ := ioutil.ReadFile("./data.txt")
  strData := strings.Split(string(file), "\n")
  data := make([]int, len(strData))
  for i, s := range strData {
    data[i], _ = strconv.Atoi(s)
  }

  depth := data[0] + data[1] + data[2]

  count := 0
  for i := 1; i < len(data) -2; i++ {
    if data[i] + data[i+1] + data[i+2] > depth {
      count++
    }
    depth = data[i] + data[i+1] + data[i+2]
  }

  // var inc uint
  // var countA, countB, countC int
  // var startB, startC int = 1, 2
  // var contC, contB, contA bool = true, true, true

  // data := []string{"199","200","208","210","200","207","240","269","260","263"}

  // for idxA, depthA := range data { 
  //   dA, err := strconv.Atoi(depthA)
  //   if err == nil && contA {
  //     // fmt.Println("A: ", dA)
  //     countA += dA
  //     if countB == 0  && contB {
  //       for cB, depthB := range data[startB:] {
  //         dB, err := strconv.Atoi(depthB)
  //         if err == nil {
  //           // fmt.Println("B: ", dB)
  //           countB += dB
  //           if countC == 0  && contC {
  //             // fmt.Println(startC)
  //             for cC, depthC := range data[startC:] {
  //               dC, err := strconv.Atoi(depthC)
  //               if err == nil {
  //                 // fmt.Println("C: ", dC)
  //                 countC += dC
  //               }
  //               if cC == 2 {
  //                 startC += 3
  //                 if (len(data) - startC) < 3 {
  //                   contC = false
  //                 }
  //                 break
  //               }
  //             }
  //             // fmt.Println("countC: ", countC)
  //           }
  //         }
  //         if cB == 2 {
  //           // fmt.Println("countB: ", countB)
  //           startB += 3
  //           if (len(data) - startB) < 3 {
  //             contB = false
  //           }
  //           break
  //         }
  //       }
  //     }
  //   }
  //   if ( ((idxA+1) % 3) == 0 ) {
  //     // fmt.Println("countA: ", countA)
  //     if countA > 0 && countB > 0 && countC > 0 {
  //     
  //       if countC > countB {
  //         inc += 1
  //         fmt.Println("countC, bigger from B", countC)
  //       }
  //       if countB > countA {
  //         inc += 1
  //         fmt.Println("countB, bigger from A", countB)
  //       } 
  //       if countA > countC {
  //         inc += 1
  //         fmt.Println("countA, bigger from C", countA)
  //       }
  //     }
  //
  //     countB = 0
  //     countA = 0
  //     countC = 0
  //   }
  //
  // }
  fmt.Println("part2 - Depth increased:", count)

}

func part1() {
  file, _ := ioutil.ReadFile("./data.txt")
  text := string(file)

  var inc uint
  var lastDepth int

  for idx, depth := range strings.Split(text, "\n") {
    v, err := strconv.Atoi(depth)
    if err == nil {
      if idx == 0 {
        lastDepth = v
        continue
      }
      if v > lastDepth {
        inc += 1
      }
      lastDepth = v
    }
  }
  fmt.Println("part1 - Depth increased:", inc)
}

func main() {
  // part1()
  part2()
}
