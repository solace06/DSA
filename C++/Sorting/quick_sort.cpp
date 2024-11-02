#include <iostream>
using namespace std;

/*
Intuition
DIVIDE AND CONQUER 
1. Pick a pivot element can be first element, last element, median or some random element
2. Place the pivot in its right position
3. Place the smaller elements on the left of the pivot and larger elements on the 
   right of the pivot

   When you place one pivot element in its right position you are left with two smaller 
   unsorted arrays.
   Then you repeat the same steps for the smaller unsorted arrays till you are left with 
   only one element in the array which is to be sorted.
Time Complexity: O(nlogn)
Space Complexity: O(1)
*/


/*
Pivot function returns the index of the pivot element after it is placed at the right
position.

what the while loop does:
    This loop executes till i and j index cross each other
    
    The fact that i and j has crossed each other indicates that all the smaller elements are 
    on the left side of the partition and the relatively larger elements are on the right side
    of the partition.

    i index:
    We iteratively use inner while loop to find the nearest element that is larger than the pivot 
    element.
    j index:
    We iteratively use inner while loop to find the nearest element that is smaller than the pivot 
    element.
    Then if i and j havent crossed each other we swap the elements placed at i and j index.

    Now when i and j index have crossed each other we swap the elements placed at low and j index to 
    place the pivot element at its right position which is at j and then return the j index or the 
    index of the pivot/partition element.
*/
int pivot(int* arr, int low, int high){
    int pivotElement=arr[low];
    int idx=low;
    int jdx=high;
    while(idx<jdx){
        while(arr[idx]<=pivotElement && idx<high)idx++;
        while(arr[jdx]>pivotElement && jdx>low)jdx--;
        if(idx<jdx){
            swap(arr[idx],arr[jdx]);
        }
    }
    swap(arr[low],arr[jdx]);
    return jdx;
}
/*
Quick Sort Algorithm:
1. Find the index/right position of the pivot element.
2. Call QuickSort function on the left unsorted array.
3. Call QuickSort on the roght unsorted array.
*/

void quickSort(int* arr, int low, int high){
    if(low<high){
        int pivotIndex=pivot(arr,low,high);
        quickSort(arr,low,pivotIndex-1);
        quickSort(arr,pivotIndex+1,high);
    }
}

void printArray(int* arr, int arrSize){
    for(int i=0; i<arrSize; i++){
        cout<<arr[i]<<" ";
    }
}

int main(){
    int arrSize;
    cout<<"Enter the size of the array\n";
    cin>>arrSize;

    int* arr= new int[arrSize];
    
    cout<<"Enter "<<arrSize<<" elements\n";
    for(int i=0; i<arrSize; i++){
        cin>>arr[i];
    }

    cout<<"Original Array\n";
    printArray(arr, arrSize);

    quickSort(arr,0,arrSize-1);

    cout<<endl;

    cout<<"Sorted Array\n";
    printArray(arr, arrSize);
}