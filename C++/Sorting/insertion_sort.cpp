#include <iostream>
using namespace std;

// Sorts array in ascending order using insertion sort
void insertionSort(int* arr, int arrSize){
    for(int i=1; i<arrSize; i++){
        int currentValue=arr[i];    // Store current element
        int idx=i-1;                // Look at previous elements
        
        // Move elements greater than currentValue one position ahead
        while(idx>=0 && arr[idx]>currentValue){
            arr[idx+1]=arr[idx];
            idx=idx-1;
        }
        arr[idx+1]=currentValue;    // Place currentValue in its correct position
    }
}

// Prints array elements with space in between
void printArray(int* arr, int arrSize){
    for(int i=0; i<arrSize; ++i){
        cout<<arr[i]<<" ";
    }
}

int main(){
    int arrSize;
    
    // Get array size from user
    cout<<"Enter the size of the array\n";
    cin>>arrSize;

    // Dynamically allocate array
    int* arr=new int[arrSize];

    // Input array elements
    cout<<"Enter "<<arrSize<<" elements\n";
    for(int i=0; i<arrSize; i++){
        cin>>arr[i];
    }

    // Display and sort
    cout<<"Original Array\n";
    printArray(arr, arrSize);

    insertionSort(arr, arrSize);

    cout<<"\n";

    cout<<"Sorted Array\n";
    printArray(arr, arrSize);

    // Free dynamically allocated memory
    delete[] arr;
    return 0;
}