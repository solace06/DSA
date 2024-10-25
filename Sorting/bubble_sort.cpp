#include<bits/stdc++.h>
using namespace std;

void printArray(int* arr, int arrSize){
    for(int i=0; i<arrSize; i++){
        cout<<arr[i]<<" ";
    }
    return;
}

void bubbleSort(int* arr, int arrSize){
    int swapped;
    for(int i=0; i<arrSize; i++){
        swapped=0;
        for(int j=i; j<arrSize-1; j++){
           
            if(arr[j]>arr[j+1]){
                swap(arr[j],arr[j+1]);
                swapped=1;
            }
        }
        if(swapped==0)break;
    }
    return;
}

int main(){
    int arrSize;
    
    cout<<"Enter the size of the array";
    cin>>arrSize;

    int* arr=new int[arrSize];

    cout<<"Enter "<<arrSize<<" elements";
    for(int i=0; i<arrSize; i++){
        cin>>arr[i];
    }

    cout<<"Unsorted Array\n";
    printArray(arr, arrSize);

    bubbleSort(arr, arrSize);
    cout<<"\n";

    cout<<"Sorted Array\n";
    printArray(arr, arrSize);

    return 0;
}