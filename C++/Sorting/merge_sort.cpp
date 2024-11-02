#include <iostream>
#include<vector>
using namespace std;

void merge(vector<int>& arr,int low, int mid, int high){
    int i=low;
    int j=mid+1;
    vector<int> temp;
    while(i<=mid && j<=high){
        if(arr[i]<=arr[j]){
            temp.push_back(arr[i]);
            i++;
        }
        else{
            temp.push_back(arr[j]);
            j++;
        }
    }
    while(j<=high){
        temp.push_back(arr[j]);
        j++;
    }
    while(i<=mid){
        temp.push_back(arr[i]);
        i++;
    }

    for(int i=low; i<=high; i++){
        arr[i]=temp[i-low];
    }
}

void mergeSort(vector<int>& arr, int low, int high){
    if(low==high)return;
    int mid=(low+high)/2;
    mergeSort(arr,low,mid);
    mergeSort(arr,mid+1,high);
    merge(arr,low,mid,high);
}

void printArray(vector<int>& arr, int size){
    for(int i=0; i<size; ++i){
        cout<<arr[i]<<" ";
    }
}

int main(){
    
    int size;

    cout<<"Enter the size of the elements\n";
    cin>>size;

    vector<int> arr(size);

    cout<<"Enter "<<size<<" elements\n";
    for(int i=0; i<size; i++){
        cin>>arr[i];
    }

    cout<<"Original Array\n";
    printArray(arr, size);

    mergeSort(arr, 0, size-1);
    cout<<endl;

    cout<<"Sorted Array\n";
    printArray(arr, size);
}