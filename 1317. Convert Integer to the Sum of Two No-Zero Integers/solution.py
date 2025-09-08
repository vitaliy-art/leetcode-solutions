class Solution:
    def getNoZeroIntegers(self, n: int) -> List[int]:
        def isNoZero(n: int) -> bool:
            return f'{n}'.rfind('0') == -1

        for i in range(1, n):
            diff = n - i
            if isNoZero(i) and isNoZero(diff):
                return [i, diff]

        return []
