/*

SAMPLE INPUT

5
1 2 3 4 5
5
1
2
3
4
5

SAMPLE OUTPUT

1
2
3
4
5


*/

package main
import ("fmt";"sort")

func binarySearch(Array []int, low int, high int, key int) int{
	// low and high are the indexes
	if low<=high{

		mid := (low+high)/2
		if Array[mid] < key{
			low = mid+1
			//fmt.Println("adding low")
			//fmt.Println(low)
			binarySearch(Array, low, high, key)
		}else if Array[mid]>key{
			high = mid-1
			//fmt.Println("Adding larger")
			//fmt.Println(high)
			binarySearch(Array, low, high, key)
		} else{
			fmt.Println(mid)
		}
	}
	return -1//key not found
}

func main() {
	// input the length of sorted sequence
	var length int
	fmt.Scanf("%d", &length)
	// create the sequence
	Array := make([]int, length+1)
	for i:=1; i<length+1; i++{
		fmt.Scanf("%d", &Array[i])
	}
	sort.Ints(Array)

	// input the query num
	var num, key int
	fmt.Scanf("%d",&num)

	for i:=0; i<num; i++ {
		fmt.Scan(&key)
		binarySearch(Array, 1, length, key)
	}
}
