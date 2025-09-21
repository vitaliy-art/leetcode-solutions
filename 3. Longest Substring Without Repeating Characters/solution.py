class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        have: set[str] = set()
        longest = 0
        current = 0
        sub_idx = -1

        for c in s:
            if c in have:
                if current > longest:
                    longest = current

                while True:
                    sub_idx += 1
                    have.remove(s[sub_idx])
                    current -= 1
                    if s[sub_idx] == c:
                        break

            current += 1
            have.add(c)

        if current > longest:
            longest = current

        return longest
