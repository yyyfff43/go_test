package algorithm

import (
	_ "fmt"
	"errors"
)

func DoSum(nameArr [] string ,scoreArr [] int) (sum int, err error)  {

	sum = 0

	for _,v := range nameArr{
		if v==""{
			err = errors.New("学生名不能为空！")
			return
		}
	}

	for _,v := range scoreArr{
		if v<0 {
			err = errors.New("请输入正整数")
		}
		sum += v;
	}

	return sum, nil
}