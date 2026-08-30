#include <bits/stdc++.h>
using namespace std;

int main() {
    int t;
    cin >> t;

    while (t--) {
        int x;
        cin >> x;

        int power = 1;

        // Find 10^number_of_digits
        while (x > 0) {
            power *= 10;
            x /= 10;
        }

        cout << power + 1 << '\n';
    }

    return 0;
}
