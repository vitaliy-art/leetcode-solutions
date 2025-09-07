function twoSum(nums: number[], target: number): number[] {
    const numsObj = {};
    const length = nums.length;
    for (let leftIdx = 0; leftIdx < length; leftIdx++) {
        const diff = target - nums[leftIdx];
        const rightIdx = numsObj[diff];
        if (rightIdx !== undefined) {
            return [leftIdx, rightIdx];
        }
        numsObj[nums[leftIdx]] = leftIdx;
    }
}
