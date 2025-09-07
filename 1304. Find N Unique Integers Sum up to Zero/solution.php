<?php

class Solution
{

    /**
     * @param Integer $n
     * @return Integer[]
     */
    function sumZero($n)
    {
        $res = [];
        $sum = 0;

        for ($idx = 0; $idx < $n - 1; $idx++) {
            $current = $idx + 1;
            $res[] = $current;
            $sum += $current;
        }

        $res[] = -$sum;
        return $res;
    }
}
