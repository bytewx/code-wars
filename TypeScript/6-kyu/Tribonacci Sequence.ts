export function tribonacci([a, b, c]: [number, number, number], n: number): number[] {
  if (n === 0) {
    return [];
  }
  
  const result = [a, b, c];
  
  for (let i = 3; i < n; i++) {
    result.push(result[i - 1] + result[i - 2] + result[i - 3]);
  }
  
  return result.slice(0, n);
}
