package slices

import "fmt"

func merge() {
    //merge 2 sorted slices
    
    //Approach 1: using extra space
    slice1:=[]int{2,4,6,8}
    slice2:=[]int{1,3,5,7}
    
    len1:=len(slice1)
    len2:=len(slice2)
    
    i:=0
    j:=0
    k:=0
    
    ans:=make([]int, len1+len2)
    for i<len1 && j<len2{
        if slice1[i]<=slice2[j]{
            ans[k]=slice1[i]
            i++
        } else{
            ans[k]=slice2[j]
            j++
        }
        k++
    }
    
    for i<len1{
        ans[k]=slice1[i]
        i++
        k++
    }
    for j<len2{
        ans[k]=slice2[j]
        j++
        k++
    }
    
    fmt.Println("Merged sorted array is:")
    fmt.Println(ans)
}