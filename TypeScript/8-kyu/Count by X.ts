export function countBy(x: number, n: number): number[] {
  const result: number[] = [];
  
  for (let i = 1; i <= n; i++) {
    result.push(x * i);
  }
  
  return result;
}
