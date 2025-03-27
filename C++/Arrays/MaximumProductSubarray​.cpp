#include <iostream>
#include <limits.h>
using namespace std;

int main() {
    int arr[]={-2,0,-1};
    int ans = INT_MIN;
    
    //brute force approach
    for(int i=0; i<3; i++){
        for(int j=i; j<3; j++){
            int curr_product=1;
            for(int k=i; k<=j; k++){
                curr_product*=arr[k];
            }
            ans=max(ans,curr_product);
        }
    }

    //better approach
    for(int i=0; i<3; i++){
        int curr_product=1;
        for(int j=i; j<3; j++){
                curr_product*=arr[j];
            ans=max(ans,curr_product);
        }
    }
    
    cout<<"Maximum product of a subarray is "<<ans;
    return 0;
}