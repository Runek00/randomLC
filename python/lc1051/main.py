from typing import List

class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        expected  = sorted(heights)
        count = 0
        for i in range(len(heights)):
            if heights[i] != expected[i]:
                count += 1
        return count

if __name__ == "__main__":
    print(Solution().heightChecker([1,1,4,2,1,3]))
    print(Solution().heightChecker([5,1,2,3,4]))
    print(Solution().heightChecker([1,2,3,4,5]))
