#include <bits/stdc++.h>
using namespace std;

int main() {
    int t;
    cin >> t;

    while (t--) {
        string s;
        cin >> s;

        int total4 = 0;
        int total2 = 0;

        // Count all 4s and 2s
        for (char ch : s) {
            if (ch == '4')
                total4++;

            if (ch == '2')
                total2++;
        }

        int prefix13 = 0;
        int answer = s.size();

        // Try every possible boundary
        for (char ch : s) {

            // At this boundary:
            // delete all 1s/3s before it
            // delete all 2s after it
            int deletions = total4 + prefix13 + total2;

            answer = min(answer, deletions);

            // Move this character to the left side
            if (ch == '1' || ch == '3') {
                prefix13++;
            }

            if (ch == '2') {
                total2--;
            }
        }

        // Check boundary after the whole string
        answer = min(answer, total4 + prefix13 + total2);

        cout << answer << '\n';
    }

    return 0;
}


