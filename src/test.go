package main
import (
            "errors"
            "fmt"
) // 我们需要使用fmt包中的Println()函数

func main() {
//      _, _, nickName := get_name_info();
//      fmt.Println(nickName);
//      fmt.Println(Addfunc(-1,2));
      var v1 int = 1
      var v2 int64 = 234
      var v3 string = "hello"
      var v4 float32 = 1.234

      MyPrintf(v1, v2, v3, v4)
}

//注意方法名小写字母开头的方法只能在本包内被调用，大写开头的才可以在包外被调用
func get_name_info()(firstName, lastName, nickName string){
      return "May", "Chan", "Chibi Maruko";
}

func Addfunc(a int, b int) (ret int, err error) {
      if a < 0 || b < 0 { // 假设这个函数只支持两个非负数字的加法
            err= errors.New("Should be non-negative numbers!")
            return
      }
      return a + b, nil // 支持多重返回值
}

func MyPrintf(args ...interface{}) {
      for _, arg := range args {
            switch arg.(type) {
            case int:
                  fmt.Println(arg, "is an int value.")
            case string:
                  fmt.Println(arg, "is a string value.")
            case int64:
                  fmt.Println(arg, "is an int64 value.")
            default:
                  fmt.Println(arg, "is an unknown type.")
            }
      }
}