/*

Ratio - TIE BREAKER

Sample Input

3 2

4 -1

4 0

10 3

5

Sample Output

3.00000000
*/

package main
import ("fmt";"math")

func find_key(Array1 []float64, Array2 []float64, K float64, low float64, high float64, R float64) {
    // calculate the pre K
    pre_k := 0.00
    R = (low + high)/2
    for i:=0; i < len(Array1); i++ {
        pre_k += Array1[i]/(Array2[i]+R)

    }

    if math.Abs( pre_k - K ) < math.Pow(10,-8) {
        fmt.Println(R)
        return
    }else if pre_k > K {
        low = R
        find_key(Array1, Array2, K , low, high, R)
    }else{
        high = R
        find_key(Array1, Array2, K , low, high, R)
    }
      

}

func main() {
    // input the number of values and column
    var query, column int
    fmt.Scanf("%d %d",&query, &column)
    // input the query nums and define two arrays
    Array1 := make([]float64, query)
    Array2 := make([]float64, query)
    
    var A, B, R float64
    
    for i:=0; i < query;i++ {
        fmt.Scanf("%f %f",&A, &B)
        Array1[i] = A
        Array2[i] = B    
    }
    var low, high float64
    
    for j:=0; j<query;j++ {
        if low > Array2[j]{
            low = Array2[j]
        }
    }
    low = -low 
    // low = 0-low
    high = 1000
    // confirm the range of R
    // input the sum key
    var K float64
    fmt.Scanf("%f",&K)
    find_key(Array1, Array2, K , low, high, R)
}

