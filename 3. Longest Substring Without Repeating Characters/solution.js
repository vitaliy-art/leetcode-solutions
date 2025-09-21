/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
    const set = new Set();
    let longest = 0;
    let current = 0;
    let subIdx = 0;

    for (let r of s) {
        if (set.has(r)) {
            if (current > longest) {
                longest = current;
            }

            for (; subIdx < s.length; subIdx++) {
                set.delete(s[subIdx]);
                current--;
                if (s[subIdx] == r) {
                    break;
                }
            }

            subIdx++;
        }

        current++;
        set.add(r);
    }

    if (current > longest) {
        longest = current;
    }

    return longest;
};
