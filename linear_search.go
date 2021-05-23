/*
  solution saves the memory
  
  
SAMPLE INPUT

5 1
1 2 3 4 1

SAMPLE OUTPUT

5


*/


package main
import (
    "fmt"
)

func main(){

    var N, M int
    // input the first line
    fmt.Scanf("%d %d", &N,&M)
    // input the second line
    // var X [100]int
    var num, index int
    for i:=1; i<=N; i++ {
        fmt.Scanf("%d", &num)
        if num == M{
            index = i
        }
    }
    if index != 0{
        fmt.Println(index)
    }else{
        // traverse the array
        fmt.Println("-1")
    }
}
