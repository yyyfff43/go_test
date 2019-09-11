package main

import "fmt"

// 我们需要使用fmt包中的Println()函数

type Integer int
func (a Integer) Less(b Integer) bool {
      return a < b
}

func main() {
      var a Integer = 1
      if a.Less(2) {
            fmt.Println(a, "Less 2")
      }
}


