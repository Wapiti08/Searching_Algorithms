/*
Will she accept him?

Sample Input

4
kalyan nobodyintheworld
rahul allgirlsallhunontheplanet
physicsguy nobodynobodylikesitbetter
lal llaggyel

Sample Output

We are only friends
Love you too
We are only friends
Love you too

*/

// Write your code here
package main
import "fmt"

func subsequence_match(string1 string, string2 string) {
    // try to find whether string1 is the subsequences of string2
    var acc, start int

    for _, sub1 := range string1 {
        // record the ord matched
        for ord2, sub2 := range string2[start:] {
            if sub1 == sub2 {
                start = ord2
                acc += 1
                break
            }
        
        }
       
    }
    if acc==len(string1){
        fmt.Println("Love you too")

    }else {
        fmt.Println("We are only friends")

    }
}


func main() {
    var cases int
    fmt.Scanf("%d", &cases)
    var string1, string2 string
    for i:=0; i<cases; i++{
        fmt.Scanf("%s %s", &string1, &string2)
        subsequence_match(string1, string2)
    }
}
