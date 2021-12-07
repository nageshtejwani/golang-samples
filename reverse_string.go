import (
"fmt"
)

func main() {
    s:="Hello World!2"
    s=reverseStr(s)
    fmt.Println(s)
}

func reverseStr(s string) string{
    r:=[]rune(s)
    x:=len(r)-1
    for y:=0;y<len(r)/2;y++{
        r[x],r[y] =r[y],r[x]
        fmt.Println(string(r))
        x--
    }
   
    return string(r)
}
