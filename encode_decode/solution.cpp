using namespace std;

#include <set>
#include <map>
#include <unordered_map>
#include <iostream>
#include <vector>
#include <utility>
#include <iomanip>
#include <iostream>
class Solution {
    public:

    
        string encode(vector<string>& strs) {
            string res = "";
            for (auto s : strs) {
                string sym = "";
                int count = 0;
                for (int i = 0; i < s.size(); i++) {
                    if (i > 0 && s[i-1] == s[i]) {
                        count++;
                    } else {
                        if (count != 0) {
                            res += sym + to_string(count);
                        }
                        count = 1;
                        sym = s[i];
                    }
                    if (i+1 == s.size()) {
                        res += sym + to_string(count);
                    }
                }
                res += string("se");
            }
            return res;
        }
    
        vector<string> decode(string s) {
            vector<string> res;
            string temp_res = "";
            for (int i = 0; i < s.size(); i+=2){
                if (s[i] == 's' && s[i+1] == 'e'){
                    res.push_back(temp_res);
                    temp_res = "";
                    continue;
                }
                for (int j = s[i+1] - '0'; j > 0; j--) 
                    temp_res += s[i];
                
            }
            return res;
        }
    };

int main(void) {
    Solution s = Solution();
    vector<string>in0 = {"#","##"};
    string res = s.encode(in0);
    cout << res << endl;
    auto d = s.decode(res);
    for (auto dec : d) {
        cout << dec << endl;
    }
    return 0;
}
    