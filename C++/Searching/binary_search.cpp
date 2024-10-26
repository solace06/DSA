#include<iostream>
using namespace std;

void binarySearch(int* arr, int arrSize, int element){
    int left=0;
    int right=arrSize-1;
    int mid;
    while(left<=right){
        //calculate the middle element of the array
        mid=(left+right)/2;
        //if the target element is found return the index
        if(arr[mid]==element){
            cout<<"Found it at "<<mid;
            return;
        }
        //if the element at mid is less than the target element
        //look in the second half of the array
        else if(arr[mid]<element){
            left=mid+1;
        }
        //else look in the first half of the array
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

    //enter the size of the array
    cout<<"Enter the size of the array";
    cin>>arrSize;

    //creating a pointer to the dynamic array
    int* arr=new int[arrSize];

    //takeing the elements of the array as input
    cout<<"Enter "<<arrSize<<" elements";
    for(int i=0; i<arrSize; ++i){
        cin>>arr[i];
    }

    //taking the element to be searched as input
    cout<<"Enter the element to be searched";
    cin>>element;

    //calling the binary search
    binarySearch(arr, arrSize, element);

    return 0;
}