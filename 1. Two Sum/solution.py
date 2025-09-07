class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numsMap = {}
        for leftIdx, left in enumerate(nums):
            diff = target - left
            if (rightIdx := numsMap.get(diff)) is not None:
                return [leftIdx, rightIdx]
            numsMap[left] = leftIdx
