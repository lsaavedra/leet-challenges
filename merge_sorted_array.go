/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.


Notes:
The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

The idea is to have 3 indexes:
1. x -> the position is the num1[m-1] element and move back
2. y -> the starting position is the num2 last element and move back
3. z -> the starting poosition is the num1 last element, moves back and overwrite the position value comparing always num1[x] vs num2[y]
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
    x := m - 1 // to check the num1[m-1] element
    y := n - 1 // to check nums2 elements
    z := m + n - 1 // move backward in num1 and overwrite the result in that pos

    for z >= 0 {
        if x < 0 { //num1 is empty
            nums1[z] = nums2[y]
            y-=1
        } else if y < 0 { //num2 is empty
            break
        } else if nums1[x] > nums2[y] {
            nums1[z] = nums1[x]
            x-=1
        } else {
            nums1[z] = nums2[y]
            y-=1
        }
        z-=1
    }
}
