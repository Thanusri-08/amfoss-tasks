#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;

    while (t--)
    {
        int n;
        cin >> n;

        vector<long long> a(n);

        for (int i = 0; i < n; i++)
        {
            cin >> a[i];
        }

        bool change = true;

        while (change)
        {
            change = false;

            for (int i = 0; i < n - 1; i++)
            {
                if (a[i] > a[i + 1])
                {
                    long long x = a[i];
                    long long y = a[i + 1];

                    a[i] = y;
                    a[i + 1] = x + y;

                    change = true;
                }
            }
        }

        cout << a[n - 1] << endl;
    }

    return 0;
}
