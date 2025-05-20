package slices

import (
    "fmt"
    "sort"
    )

func sorting() {
    //sorting slices using in built functions
    org:=[]int{3,6,4,1,2,5,3}
    fmt.Printf("original slice: %v", org)
    fmt.Println()
    
    //ascending order
    asc:=make([]int,len(org))
    copy(asc,org)
    sort.Ints(asc)
    fmt.Printf("slice in ascending order: %v", asc)
    fmt.Println()
    
     //descending order
    desc:=make([]int,len(org))
    copy(desc,org)
    sort.Sort(sort.Reverse(sort.IntSlice(desc)))
    fmt.Printf("slice in descending order: %v", desc)
    fmt.Println()
}