#include <iostream>
using namespace std;

void mergeSort(int* arr, int size){

}

void printArray(int* arr, int size){
    for(int i=0; i<size; ++i){
        cout<<arr[i]<<" ";
    }
}

int main(){
    
    int size;

    cout<<"Enter the size of the elements\n";
    cin>>size;

    int* arr= new int[size];

    cout<<"Enter "<<size<<" elements\n";
    for(int i=0; i<size; i++){
        cin>>arr[i];
    }

    cout<<"Original Array\n";
    printArray(arr, size);

    mergeSort(arr, size);
    cout<<endl;

    cout<<"Sorted Array\n";
    printArray(arr, size);
}