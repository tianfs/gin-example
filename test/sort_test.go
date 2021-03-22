package test

import (
    "fmt"
    "math/rand"
    "sync"
    "testing"
)

var wg sync.WaitGroup

func TestSort(test *testing.T) {

    nums := []int{}
    for i := 0; i <= 20; i++ {
        nums = append(nums, rand.Intn(100))
    }

    newNums := doKuaipai(nums)

    fmt.Println("排序", nums, len(nums))
    fmt.Println("排序x新", newNums, len(newNums))
}

func doKuaipai(nums []int) []int {

    if len(nums) <= 1 {
        return nums
    }
    flat := nums[0]
    lefts := []int{}
    rights := []int{}
    for _, v := range nums {
        if v < flat {
            lefts = append(lefts, v)
        }
        if v > flat {
            rights = append(rights, v)
        }
    }
    fmt.Println("排序左右", lefts, rights)

    lefts = doKuaipai(lefts)
    rights = doKuaipai(rights)

    newNums := []int{}
    newNums = append(newNums, lefts...)
    newNums = append(newNums, flat)
    newNums = append(newNums, rights...)

    return newNums

}
