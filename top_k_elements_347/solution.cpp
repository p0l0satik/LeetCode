using namespace std;

#include <set>
#include <map>
#include <unordered_map>
#include <iostream>
#include <vector>

class Solution {
    public:
        vector<int> topKFrequent(vector<int>& nums, int k) {
            unordered_map<int, int> freq;
            vector<vector<int>> buckets (nums.size(), vector<int>(0));
            for (int num : nums) {
                freq[num]++;
            }
            for(pair<int, int> k_v : freq) {
                buckets[k_v.second-1].push_back(k_v.first);
            }
            vector<int> res;
            for (int i = buckets.size()-1; res.size() < k; i--){
                if (buckets[i].size() > 0) {
                    res.insert(res.end(), buckets[i].begin(), buckets[i].end());
                }
            }
            return res;
        }
    };
    
int main(void) {
    Solution s = Solution();
    vector<int>in0 = {7,7};
    cout << s.topKFrequent(in0, 1)[0] << endl;
    return 0;
}