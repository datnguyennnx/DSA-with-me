/*
You are given a rectangular board of M × N squares. Also you are given an unlimited number of standard domino pieces of 2 × 1 squares. You are allowed to rotate the pieces. You are asked to place as many dominoes as possible on the board so as to meet the following conditions:

1. Each domino completely covers two squares.

2. No two dominoes overlap.

3. Each domino lies entirely inside the board. It is allowed to touch the edges of the board.

Find the maximum number of dominoes, which can be placed under these restrictions.

Input
In a single line you are given two integers M and N — board sizes in squares (1 ≤ M ≤ N ≤ 16).

Output
Output one number — the maximal number of dominoes, which can be placed.
*/
#include <stdio.h>
#include <math.h>
void checkRow(int nums1, int nums2){
        if(nums1 <= 16 && nums2 <= 16){
             printf("%d", (nums1 *nums2)/2);
        }
    }
int main (){
    int nums1, nums2;
    scanf("%d", &nums1);
    scanf("%d", &nums2);
    checkRow(nums1, nums2);
    return 0;
}