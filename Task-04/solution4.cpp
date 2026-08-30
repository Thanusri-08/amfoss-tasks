#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;

    vector<int> prime;

    for (int i = 2; i <= 110000; i++)
    {
        bool ok = true;

        for (int j = 2; j * j <= i; j++)
        {
            if (i % j == 0)
            {
                ok = false;
                break;
            }
        }

        if (ok)
            prime.push_back(i);
    }

    while (t--)
    {
        int n;
        cin >> n;

        for (int i = 0; i < n; i++)
        {
            long long x = 1LL * prime[i] * prime[i + 1];
            cout << x << " ";
        }

        cout << endl;
    }

    return 0;
}
