#include <iostream>
#include <bits/stdc++.h>
#include <set>
using namespace std;

int main() {
    int arr[6]={-1,0,1,2,-1,-4};
    set <vector<int>> ans;
    for(int i=0; i<6; i++){
        for(int j=i+1; j<6; j++){
            for(int k=j+1; k<6; k++){
                if(arr[i]+arr[j]+arr[k]==0){
                    vector<int> temp = {arr[i],arr[j],arr[k]};
                    sort(temp.begin(), temp.end());
                    ans.insert(temp);
                }
            }
        }
    }
    
    for(vector<int> i: ans){
        cout<<"{";
        for(int j: i){
            cout<<j<<",";
        }
        cout<<"}";
    }
    return 0;
}