/**
 * @param {number} n
 * @return {number[]}
 */
var getNoZeroIntegers = function (n) {
    const isNoZero = (n) => !`${n}`.includes("0");

    for (let i = 1; i < n; i++) {
        const diff = n - i;
        if (isNoZero(i) && isNoZero(diff)) {
            return [i, diff];
        }
    }
};
