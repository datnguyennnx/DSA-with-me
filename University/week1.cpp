#include <iostream>
#include <cstdlib>

using namespace std;

int main (){
	long long n;
	cout << "Nhap n: ";
	cin >> n;
	long long k;
	cout << "Nhap k: ";
	cin >> k;

	int a[n] = {};
	
	long long result = 0;
	for ( int i = 0; i < n; i++){
		a[i] = rand() % 5 + 1;
		cout << a[i] << " ";
	}
	cout << endl;
	

	for ( int i = 0; i < k; i++){
		int p = rand() % n % 1000;
		int q = rand() % n + p;
		if ( q > n) {
			q = n; 
		} else {
			q = q;
		}
		cout << "p: " << p << endl;
		cout << "q: " << q << endl;
		for ( int i = p; i < q; i++){
			 result += a[i];
		}
	}

	cout << "result: " << result;
	return 0;

}
