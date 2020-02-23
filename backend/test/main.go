package main

import (
    "fmt"
    "testgo/models"
)

type people struct {
    name string
}
func main(){
    r3 := returnNil3()
    r4 := returnNil4()
    fmt.Println(r3, "-------", r4)
    if r3 != nil{
        fmt.Println("r3 != nil")
    }else {
        fmt.Println("r3 == nil")
    }
    if r4 != nil{
        fmt.Println("r4 equal nil is true")
    }else {
        fmt.Println("r4 equal nil is not true")
    }

}


func returnNil3()(r *models.Project){
    return nil
}
func returnNil4()(r *[]models.Project){
    return nil
}


