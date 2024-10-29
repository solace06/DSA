#include<iostream>
using namespace std;

//function to print the array
void printArray(int* arr, int arrSize){
    for(int i=0; i<arrSize; ++i){
        cout<<arr[i]<<" ";
    }
}

//function to sort the array
/*
Algorithm:
Find the i th minimum element from the array
and place it at i th position
Track the index of the i th minimum element 
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
void selectionSort(int* arr, int arrSize){
    int minElementIndex;
    for(int i=0; i<arrSize-1; i++){
        minElementIndex=i;
        for(int j=i+1; j<arrSize; ++j){
            if(arr[minElementIndex]>arr[j]){
                minElementIndex=j;
            }
        }
        swap(arr[minElementIndex],arr[i]);
    }
}

int main(){
    int arrSize;
    cout<<"Enter the size of the array\n";
    cin>>arrSize;

    //creating pointer to the dynamic array
    int* arr=new int[arrSize];

    cout<<"Enter "<<arrSize<<" elements\n";
    for(int i=0; i<arrSize; ++i){
        cin>>arr[i];
    }

    cout<<"Original Array\n";
    printArray(arr, arrSize);

    selectionSort(arr, arrSize);

    cout<<"\n";

    cout<<"Sorted Array\n";
    printArray(arr, arrSize);

    //Free dynamically allocated memory
    delete []arr;
    return 0;
}