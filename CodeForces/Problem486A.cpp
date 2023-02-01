/*
For a positive integer n let's define a function f:

f(n) =  - 1 + 2 - 3 + .. + ( - 1)nn

Your task is to calculate f(n) for a given integer n.

Input
The single line contains the positive integer n (1 ≤ n ≤ 1015).

Output
Print f(n) in a single line.
*/

#include <iostream>
#include <cmath>

using namespace std;

int main(){
	long long nums;
	cin >> nums;
	if ( nums % 2 == 0){
		cout << nums/2;
	} else {
		cout << (-nums - 1)/2;
	}
	return 0;
}
