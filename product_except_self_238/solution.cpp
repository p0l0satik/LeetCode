using namespace std;

#include <set>
#include <iostream>
#include <vector>


class Solution {
    public:
        vector<int> productExceptSelf(vector<int>& nums) {
            vector<int> prefix;
            vector<int> suffix(nums.size(), 0);
            int cum_summ = 1;

            for (auto n : nums) {
                prefix.push_back(cum_summ);
                cum_summ *= n;
            }
            cum_summ = 1;
            for (int i = nums.size() -1; i >= 0; i--){
                suffix[i] = cum_summ;
                cum_summ *= nums[i];
            }

            vector<int> res;
            for (int i = 0; i < nums.size(); i++){
                res.push_back(prefix[i]*suffix[i]);
            }
            return res;
    }
};

int main(void) {
    Solution s = Solution();
    vector<int> in {1, 2, 4, 6}; 
    s.productExceptSelf(in);
    return 0;
}