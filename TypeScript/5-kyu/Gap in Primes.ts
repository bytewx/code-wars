export const gap = (g: number, m: number, n: number): number[] | null => {
    const isPrime = (num: number): boolean => {
        if (num < 2) return false;
        for (let i = 2; i * i <= num; i++) {
            if (num % i === 0) return false;
        }
        return true;
    };

    let previousPrime = null;

    for (let i = m; i <= n; i++) {
        if (isPrime(i)) {
            if (previousPrime !== null && i - previousPrime === g) {
                return [previousPrime, i];
            }

            previousPrime = i;
        }
    }
    return null;
};
