#include <iostream>
using namespace std;

int main() {
    int arr[]={-1,1,0,-3,3};
    int ans[5];
    int prefix[5];
    int suffix[5];
    
    //brute force approach
    for(int i=0; i<5; i++){
        int curr_product=1;
        for(int j=0; j<5; j++){
            if(i!=j) curr_product*=arr[j];
        }
        ans[i]=curr_product;
    }

    //better approach
    for(int i=0; i<5; i++){
        if(i==0)prefix[i]=1;
        else prefix[i] = prefix[i-1]*arr[i-1];
    }
    
    for(int i=4; i>0; i--){
        if(i==4)suffix[i]=1;
        else suffix[i] = suffix[i+1]*arr[i+1];
    }
    
    for(int i=0; i<5; i++){
        ans[i]=prefix[i]*suffix[i];
    }
    
    cout<<"The output array is as follows "<<endl;
    for(int i=0; i<5; i++){
        cout<<ans[i]<<" ";
    }
    return 0;
}