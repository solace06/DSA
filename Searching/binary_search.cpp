#include<iostream>
using namespace std;

void binarySearch(int* arr, int arrSize, int element){
    int left=0;
    int right=arrSize-1;
    int mid;
    while(left<=right){
        mid=(left+right)/2;
        if(arr[mid]==element){
            cout<<"Found it at "<<mid;
            return;
        }
        else if(arr[mid]<element){
            left=mid+1;
        }
        else{
            right=mid-1;
        }
    }
    cout<<"Not Found!";
    return;
}

int main(){
    
    int arrSize;
    int element;

    cout<<"Enter the size of the array";
    cin>>arrSize;

    int* arr=new int[arrSize];

    cout<<"Enter "<<arrSize<<" elements";
    for(int i=0; i<arrSize; ++i){
        cin>>arr[i];
    }

    cout<<"Enter the element to be searched";
    cin>>element;

    binarySearch(arr, arrSize, element);

    return 0;
}