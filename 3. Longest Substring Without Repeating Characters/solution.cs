public class Solution
{
    public int LengthOfLongestSubstring(string s)
    {
        var set = new HashSet<char>();
        var longest = 0;
        var current = 0;
        var subIdx = 0;

        foreach (var r in s)
        {
            if (set.Contains(r))
            {
                if (current > longest)
                {
                    longest = current;
                }

                for (; subIdx < s.Length; subIdx++)
                {
                    set.Remove(s[subIdx]);
                    current--;
                    if (s[subIdx] == r)
                    {
                        break;
                    }
                }

                subIdx++;
            }

            current++;
            set.Add(r);
        }

        if (current > longest)
        {
            longest = current;
        }

        return longest;
    }
}
