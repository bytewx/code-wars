export function getMiddle(s: string): string {
    const len = s.length;
    const mid = Math.floor(len / 2);

    return len % 2 === 0 ? s[mid - 1] + s[mid] : s[mid];
}
