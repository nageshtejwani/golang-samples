package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f []int 
	a:=[]int{1,2,3,4,5} //initialization with elements
	b:=make([]int,3) //initialization with length 3
	c:=make([]int,3,100)  //initialization with length 3 and 100 capacity
	d:=[3]int{1,2,3}//array 
	f=d[0:1] //slice from array
	fmt.Printf("Slice a %v  :: length  %d  ,capacity ::%d \n",a ,len(a) ,cap(a))
	fmt.Printf("Slice b %v  :: length  %d  ,capacity ::%d \n",b ,len(b) ,cap(b))
	fmt.Printf("Slice c %v  :: length  %d  ,capacity ::%d \n",c ,len(c) ,cap(c))
	fmt.Printf("Slice f %v  :: length  %d  ,capacity ::%d \n",f ,len(f) ,cap(f))
 
}
