#include<iostream>
using namespace std;

int main(){
    int arrSize;
    int element;

    //taking the size of the array as input
    cout<<"Enter the size of the array";
    cin>>arrSize;

    //creating a pointer to the array
    int* arr= new int[arrSize];

    //Taking the elements of the array as input
    cout<<"Enter "<<arrSize<<" elements";
    for(int i=0; i<arrSize; ++i){
        cin>>arr[i];
    }
    
    //taking the input(element to be searched for) from the user
    cout<<"Enter the element to be searched\n";
    cin>>element;

    //iterating through the array to find the element
    for(int i=0; i<arrSize; ++i){
        if(arr[i]==element){
            cout<<"\nFound it!!";
            //freeing the allocated space
            delete[] arr;
            return 0;
        }
    }
    cout<<"\nOops better luck next time";
    delete[] arr;
    return 0;

    //Time Complexity:O(n)
    //Space Complexity:O(1)
}