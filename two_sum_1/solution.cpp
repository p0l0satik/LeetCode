using namespace std;
#include <map>
#include <vector>

class Solution {
    public:
        vector<int> twoSum(vector<int>& nums, int target) {
            map<int,int> subs;
    
            for (int i = 0; i < nums.size(); i++) {
                subs[target - nums[i]] = i;
            }
            for (int i = 0; i < nums.size(); i++){
                if (subs.find(nums[i]) != subs.end()) {
                    int ii = subs[nums[i]];
                    if (ii == i) {
                        continue;
                    }
                    if (ii > i) {
                        return vector<int>{i, ii};
                    } else {
                        return vector<int>{ii, i};
                    }
                } 
            }
        }
    };
    