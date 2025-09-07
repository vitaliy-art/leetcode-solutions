public class Solution {
    public int[] SumZero(int n) {
        var res = new int[n];
        var sum = 0;

        for (var idx = 0; idx < n - 1; idx++) {
            var current = idx + 1;
            res[idx] = current;
            sum += current;
        }

        res[n-1] = -sum;
        return res;
    }
}
