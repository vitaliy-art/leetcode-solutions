<?php

class Solution
{

    /**
     * @param Integer $n
     * @return Integer[]
     */
    function getNoZeroIntegers($n)
    {
        $isNoZero = fn(int $n) => !str_contains("$n", '0');

        for ($i = 1; $i < $n; $i++) {
            $diff = $n - $i;
            if ($isNoZero($i) && $isNoZero($diff)) {
                return [$i, $diff];
            }
        }
    }
}
