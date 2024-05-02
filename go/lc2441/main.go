package main

func findMaxK(nums []int) int {
 out := -1
 numsMap := make(map[int]bool)
 for _, n := range nums {
     _, ok := numsMap[-n]
     if ok {
         if n > out {
             out = n
         } else if -n > out {
             out = -n
         }
     }
     numsMap[n] = true

 }
 return out
}

func main() {
 } 