using namespace std;

#include <string>
#include <map>
#include <iostream>


class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.length() != t.length()) {
            return false;
        }
        map<char, int> m_s{}, m_t{};
        for (int i = 0; i < s.length();i++) {
            m_s[s[i]]++;
            m_t[t[i]]++;    
        }
        return m_s == m_t;
    }
};

int main(void) {
    Solution s = Solution();
    string s1 = "anagram", t1 = "nagaram";
    string s2 = "rat", t2 = "car";
    cout << s.isAnagram(s1, t1) << endl;
    cout << s.isAnagram(s2, t2) << endl;
    return 0;
}