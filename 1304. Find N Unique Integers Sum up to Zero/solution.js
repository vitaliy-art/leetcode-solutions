/**
 * @param {number} n
 * @return {number[]}
 */
var sumZero = function (n) {
    const res = [];
    let sum = 0;

    for (let idx = 0; idx < n - 1; idx++) {
        const current = idx + 1;
        res.push(current);
        sum += current;
    }

    res.push(-sum);
    return res;
};
