#include <iostream>
using namespace std;

void quickSort(int* arr, int arrSize){
    
}

void printArray(int* arr, int arrSize){
    for(int i=0; i<arrSize; i++){
        cout<<arr[i]<<" ";
    }
}

int main(){
    int arrSize;
    cout<<"Enter the size of the array";
    cin>>arrSize;

    int* arr= new int[arrSize];
    
    cout<<"Enter "<<arrSize<<" elements";
    for(int i=0; i<arrSize; i++){
        cin>>arr[i];
    }

    cout<<"Original Array";
    printArray(arr, arrSize);

    quickSort(arr, arrSize);

    cout<<"Sorted Array";
    printArray(arr, arrSize);
}