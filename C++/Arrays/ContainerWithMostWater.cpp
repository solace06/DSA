#include <iostream>
using namespace std;

int main() {
    int height[] = {1,8,6,2,5,4,8,3,7};
    int n = sizeof(height) / sizeof(height[0]); 
    int ans = 0;

   
   //brute force approach
    for(int i = 0; i < n; i++) {
        for(int j = i + 1; j < n; j++) {
            int h = min(height[i], height[j]);
            int b = j - i;
            ans = max(ans, h * b);
        }
    }

    cout << "Maximum amount of water a container can store: " << ans;
    return 0;
}
