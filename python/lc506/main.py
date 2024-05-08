from typing import List

class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
        place = score.copy()
        place.sort(reverse = True)
        unplace = {place[i]: i for i in range(len(place))}
        out = []
        rank = ["Gold Medal", "Silver Medal", "Bronze Medal"]
        for i in range(len(score)):
            if score[i] == place[0]:
                out.append("Gold Medal")
            elif score[i] == place[1]:
                out.append("Silver Medal")
            elif score[i] == place[2]:
                out.append("Bronze Medal")
            else:
                out.append(str(unplace[score[i]] + 1))
        return out 

if __name__ == "__main__":
    score = [5, 4, 3, 2, 1]
    print(Solution().findRelativeRanks(score))
    score = [10, 3, 8, 9, 4]
    print(Solution().findRelativeRanks(score))
