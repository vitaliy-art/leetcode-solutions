<?php

class Solution
{

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target)
    {
        $numsMap = [];
        foreach ($nums as $leftIdx => $left) {
            $diff = $target - $left;
            if (($rightIdx = $numsMap[$diff] ?? null) !== null) {
                return [$leftIdx, $rightIdx];
            }
            $numsMap[$left] = $leftIdx;
        }
    }
}
