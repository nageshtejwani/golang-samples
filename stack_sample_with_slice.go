func main() {
    a:=[]int{1,2,3,4,5}
    fmt.Println("List before any operations ::")
    fmt.Println(a)
    fmt.Println("After push ::")
    aList:=push(a,2)
    fmt.Println(aList)
    fmt.Println("After pop ::")
    aList=pop(aList)
    fmt.Println(aList)
}

func push(myList []int , a int  )[]int{
    aList:= make([]int,1)
    aList[0]= a
    aList = append(aList,myList...)
    return aList
}


func pop(myList []int )[]int{
    return myList[1:]
}
