#include <bits/stdc++.h> 
using namespace std;

void jumpSearch(int*arr, int arrSize, int element){
    int stepSize=sqrt(arrSize);
    int currIndex=0;

    while(currIndex<arrSize){
        if(arr[currIndex]==element){
            cout<<"Found it!";
            return;
        }
        else if(arr[currIndex]>element){
            break;
        }
        else{
            currIndex+=stepSize;
        }
    }

    int startIndex=currIndex-stepSize;
    if(startIndex<0)startIndex=0;

    for(int i=startIndex; i<=currIndex; ++i){
        if(arr[i]==element){
            cout<<"Found it!";
            return;
        }
    }
    cout<<"Not Found";
    return;
}


int main(){
    int arrSize;
    int element;

    cout<<"Enter the size of the element";
    cin>>arrSize;

    int* arr=new int[arrSize];
    cout<<"Enter "<<arrSize<<" elements";
    for(int i=0; i<arrSize; i++){
        cin>>arr[i];
    }

    cout<<"Enter the element to be searched for";
    cin>>element;

    jumpSearch(arr, arrSize, element);
    return 0;
}