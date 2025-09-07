/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
    const numsMap = new Map();
    for (let leftIdx = 0; leftIdx < nums.length; leftIdx++) {
        left = nums[leftIdx];
        const diff = target - left;
        const rightIdx = numsMap.get(diff);
        if (rightIdx !== undefined) {
            return [leftIdx, rightIdx];
        }
        numsMap.set(left, leftIdx);
    }
};
