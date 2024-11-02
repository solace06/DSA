#include <iostream>
using namespace std;

/*
Intuition
1. Pick a pivot element can be first element, last element, median or some random element
2. Place the pivot in its right position
3. Place the smaller elements on the left of the pivot and larger elements on the 
   right of the pivot

   When you place one pivot element in its right position you are left with two smaller 
   unsorted arrays.
   Then you repeat the same steps for the smaller unsorted arrays till you are left with 
   only one element in the array which is to be sorted.
Time Complexity: O()
Space Complexity: O()
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