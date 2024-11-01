export function positiveSum(arr:number[]):number {
  return arr.reduce((sum, value) => {
    return value > 0 ? sum + value : sum;
  }, 0);
}
