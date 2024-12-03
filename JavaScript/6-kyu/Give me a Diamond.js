function diamond(n) {
    if (n <= 0 || n % 2 === 0) return null;
    
    let result = '';
    const mid = Math.floor(n / 2);

    for (let i = 0; i < n; i++) {
        const stars = n - Math.abs(mid - i) * 2;
        const spaces = Math.abs(mid - i);
        result += ' '.repeat(spaces) + '*'.repeat(stars) + '\n';
    }

    return result;
}
