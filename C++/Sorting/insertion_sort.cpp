#include <iostream>
using namespace std;

void insertionSort(int* arr, int arrSize){
    for(int i=1; i<arrSize; i++){
        int currentValue=arr[i];
        int idx=i-1;
        while(idx>=0 && arr[idx]>currentValue){
            arr[idx+1]=arr[idx];
            idx=idx-1;
        }
        arr[idx+1]=currentValue;
    }
}


void printArray(int* arr, int arrSize){
    for(int i=0; i<arrSize; ++i){
        cout<<arr[i]<<" ";
    }
}


int main(){

    int arrSize;

    cout<<"Enter the size of the array\n";
    cin>>arrSize;

    int* arr=new int[arrSize];

    cout<<"Enter "<<arrSize<<" elements\n";
    for(int i=0; i<arrSize; i++){
        cin>>arr[i];
    }

    cout<<"Original Array\n";
    printArray(arr, arrSize);

    insertionSort(arr, arrSize);

    cout<<"\n";

    cout<<"Sorted Array\n";
    printArray(arr, arrSize);
}