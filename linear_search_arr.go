/*

SAMPLE INPUT

5 1
1 2 3 4 1

SAMPLE OUTPUT

5

*/

package main
import (
    "fmt"
    "os"
)

func main(){

    var N, M int
    // input the first line
    fmt.Scanf("%d %d", &N,&M)
    // input the second line
    X := make([]int, N+1)
    // var X [100]int
    for i:=1; i<=N; i++ {
        fmt.Scanf("%d", &X[i])
    }

    // traverse the array
    for ord:= N; ord>=1; ord-- {
        if X[ord] == M {
            fmt.Println(ord)
            os.Exit(3)
            }
        
    }
    fmt.Println("-1")
    
}
