/*
  Problem: Two Sum
  url: https://leetcode.com/problems/two-sum/description/

  Description: Given an array of integers, return indices of the two numbers such that they add up to a specific target.
  You may assume that each input would have exactly one solution, and you may not use the same element twice.

  Example.
    Given nums = [2, 7, 11, 15], target = 9,
    Because nums[0] + nums[1] = 2 + 7 = 9,
    return [0, 1].
*/

package main
import "fmt"

func main() {
  nums := []int {2, 7, 11, 15}
  fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
  var indices []int
  length := len(nums)
  for i := 0; i< length-1; i++ {
    for j := i+1; j< length; j++ {
      if (nums[i] + nums[j] == target) {
        indices = append(indices, i, j)
      }
    }
  }
  return indices
}

/*******************************

OUTPUT:

[0 1]

*******************************/