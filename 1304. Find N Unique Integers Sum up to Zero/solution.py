class Solution:
    def sumZero(self, n: int) -> List[int]:
        if n == 1:
            return [0]

        res = []
        sum = 0
        for idx in range(n-1):
            cur = idx + 1
            res.append(cur)
            sum += cur

        res.append(-sum)
        return res
