package main
import (
      "fmt"
      "reflect"
)
//利用反射reflect机制遍历结构体

type person struct {
      name string
      age  int
}
func main() {
      v := reflect.ValueOf(person{"steve", 30})
      count := v.NumField()
      for i := 0; i < count; i++ {
            f := v.Field(i)
            switch f.Kind() {
            case reflect.String:
                  fmt.Println(f.String())
            case reflect.Int:
                  fmt.Println(f.Int())
            }
      }
}