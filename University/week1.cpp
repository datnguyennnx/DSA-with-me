/*
	Input n for generate Array[n] = random 1 -> 5
	Input k for loop calculate sum index from p -> q ( p and q is random for each k loop)
	Condition: 0 < p < 10^3 and n < 10^8
	Testcase if n > 10^8 what if happen?
*/
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
