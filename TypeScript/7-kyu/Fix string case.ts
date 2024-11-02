export function solve(s: string): string {
  let lowerCount = 0;
  let upperCount = 0;

  for (let char of s) {
    if (char === char.toLowerCase()) {
      lowerCount++;
    } else {
      upperCount++;
    }
  }

  return lowerCount >= upperCount ? s.toLowerCase() : s.toUpperCase();
}
