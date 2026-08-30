#include <bits/stdc++.h>
using namespace std;

int main() {
    int t;
    cin >> t;

    while (t--) {
        int n, c;
        cin >> n >> c;

        vector<int> a(n), b(n);

        for (int i = 0; i < n; i++) {
            cin >> a[i];
        }

        for (int i = 0; i < n; i++) {
            cin >> b[i];
        }

        // Step 1: Try without rearranging
        bool possible = true;
        int cost = 0;

        for (int i = 0; i < n; i++) {
            if (a[i] < b[i]) {
                possible = false;
                break;
            }

            cost += a[i] - b[i];
        }

        if (possible) {
            cout << cost << '\n';
            continue;
        }

        // Step 2: Try after rearranging
        sort(a.begin(), a.end());
        sort(b.begin(), b.end());

        possible = true;
        cost = c;

        for (int i = 0; i < n; i++) {
            if (a[i] < b[i]) {
                possible = false;
                break;
            }

            cost += a[i] - b[i];
        }

        if (possible) {
            cout << cost << '\n';
        } else {
            cout << -1 << '\n';
        }
    }

    return 0;
}
