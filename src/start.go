package main

import (
      "fmt"
      "algorithm"
)

func main() {
      sNameArr := []string{"张三","李四","王五","赵六","田七"}
      sScoreArr := []int{96,98,100,95,95}
      fmt.Println("学生成绩列表",sNameArr,sScoreArr)
      sSum , _:= algorithm.DoSum(sNameArr,sScoreArr)
      fmt.Printf("班级总分为%d",sSum)
      fmt.Println("")
      sAvg , _:= algorithm.DoAverage(sNameArr,sScoreArr)
      fmt.Printf("班级平均分为%f",sAvg)
}
