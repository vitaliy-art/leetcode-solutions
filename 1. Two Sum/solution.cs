public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        var numsDict = new Dictionary<int, int>();
        for (var leftIdx = 0; leftIdx < nums.Length; leftIdx++) {
            var left = nums[leftIdx];
            var diff = target - left;
            if (numsDict.TryGetValue(diff, out var rightIdx)) {
                return new int[]{leftIdx, rightIdx};
            }
            numsDict[left] = leftIdx;
        }
        return new int[]{};
    }
}
