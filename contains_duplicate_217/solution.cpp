using namespace std;

#include <set>
#include <iostream>
#include <vector>

class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        set<int> saw;
        for (int i : nums) {
            if (saw.contains(i)) {
                return true;
            } 
            saw.insert(i);
        }
        return false;
    }
};

int main(void){
    Solution s = Solution();
    vector<int> v1 = {1,2,3,1};
    vector<int> v2 = {1,2,3,4};
    vector<int> v3 = {1,1,1,3,3,4,3,2,4,2};

    cout << s.hasDuplicate(v1) << endl;
    cout << s.hasDuplicate(v2) << endl;
    cout << s.hasDuplicate(v3) << endl;

}