#include<bits/stdc++.h>
using namespace std;

/**
 * Prints the elements of an array with proper formatting
 * @param arr Pointer to the array to be printed
 * @param arrSize Size of the array
 * Time Complexity: O(n) where n is the size of array
 * Space Complexity: O(1)
 */
void printArray(int* arr, int arrSize){
    for(int i=0; i<arrSize; i++){
        cout<<arr[i]<<" ";
    }
    return;
}

/**
 * Implements the Bubble Sort algorithm to sort an array in ascending order
 * @param arr Pointer to the array to be sorted
 * @param arrSize Size of the array
 * 
 * Algorithm:
 * 1. Uses two nested loops:
 *    - Outer loop (i) represents the number of passes
 *    - Inner loop (j) performs comparisons and swaps in each pass
 * 2. In each iteration, compares adjacent elements and swaps if they are in wrong order
 * 3. Optimization: Uses 'swapped' flag to detect if array is already sorted
 * 
 * Time Complexity: 
 * - Best Case: O(n) when array is already sorted
 * - Average and Worst Case: O(n²) where n is the size of array
 * Space Complexity: O(1) as it sorts in-place
 */
void bubbleSort(int* arr, int arrSize){
    int swapped;
    // Each pass will place the largest remaining element at the end
    for(int i=0; i<arrSize-1; i++){
        swapped=0;
        // Compare adjacent elements and swap if they are in wrong order
        // Note: arrSize-i-1 because last i elements are already sorted
        for(int j=0; j<arrSize-i-1; j++){
            if(arr[j]>arr[j+1]){
                swap(arr[j],arr[j+1]);
                swapped=1;
            }
        }
        // If no swapping occurred, array is already sorted
        if(swapped==0)break;
    }
    return;
}

int main(){
    int arrSize;
    
    //Take the size of the array as input
    cout<<"Enter the size of the array";
    cin>>arrSize;

    //creating pointer to the dynamic array
    int* arr=new int[arrSize];

    //taking the elements of the array as input
    cout<<"Enter "<<arrSize<<" elements";
    for(int i=0; i<arrSize; i++){
        cin>>arr[i];
    }

    //printing the unsorted array
    cout<<"Unsorted Array\n";
    printArray(arr, arrSize);

    //Caaling the bublle sort function
    bubbleSort(arr, arrSize);
    cout<<"\n";

    //Printing the sorted array
    cout<<"Sorted Array\n";
    printArray(arr, arrSize);

     // Free dynamically allocated memory
    delete[] arr;
    return 0;
}