package algorithm

import "errors"

func DoAverage(nameArr [] string ,scoreArr [] int)(avgvalue float32,err error){

	sum := 0;

    for _,v := range scoreArr{
    	if v<0 {
    		err = errors.New("分数必须为正整数")
    		return
		}
    	sum += v;
	}

//	panic(len(nameArr))
	avgScore := sum/len(nameArr)
	return float32(avgScore), nil;
}
