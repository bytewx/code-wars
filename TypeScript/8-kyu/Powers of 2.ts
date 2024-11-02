export function powersOfTwo(n:number):number[]{
  const result: number[] = [];
  
  for (let i = 0; i <= n; i++) {
    result.push(2 ** i);
  }
  
  return result;
}
