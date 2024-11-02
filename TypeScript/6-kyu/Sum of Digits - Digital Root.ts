export const digitalRoot = (n: number): number => {
    if (n < 10) {
        return n;
    }

    const sum = n.toString().split('').reduce((acc, digit) => acc + Number(digit), 0);

    return digitalRoot(sum);
};
