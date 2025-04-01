using namespace std;

#include <set>
#include <map>
#include <unordered_map>
#include <iostream>
#include <vector>

class Solution {
    public:
        vector<vector<string>> groupAnagrams(vector<string>& strs) {
            unordered_map<string,vector<string>> anagramsSet;
            for (string sss : strs) {
                vector<int> prehash(26*2, 0); 
                for (char c : sss) {
                    prehash[c - 'a']++;
                }
                string hash = "";
                for (int i = 0; i < prehash.size(); i++){
                    if (prehash[i] > 0) {
                        hash += string(1, 'a'+i)+string(1,prehash[i]);
                    }
                }
                if (anagramsSet.find(hash) != anagramsSet.end()) {
                    anagramsSet[hash].push_back(sss);
                } else {
                    anagramsSet[hash] = vector<string> {sss};
                }
            }
            vector<vector<string>> res;
            for (auto const& [key, val] : anagramsSet){
                res.push_back(val);
            }
    
            return res;
        }
    };

    int main(void) {
        Solution s = Solution();
        vector<string>in0 = {"act","pots","tops","cat","stop","hat"};
        vector<string>in1 = {"ddddddddddg","dgggggggggg"};
        s.groupAnagrams(in0);
        return 0;
    }