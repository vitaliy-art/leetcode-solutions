public class Solution {
    private bool IsNoZero(int n) => !$"{n}".Contains("0");

    public int[] GetNoZeroIntegers(int n) {
        for (var i = 1; i < n; i++) {
            var diff = n - i;
            if (this.IsNoZero(i) && this.IsNoZero(diff)) {
                return new int[]{i, diff};
            }
        }

        return new int[]{};
    }
}
