export function sumArray(array:number[] | null) : number {
  if (!array || array.length <= 1) {
    return 0;
  }
  
  const min = Math.min(...array);
  const max = Math.max(...array);
  
  const totalSum = array.reduce((sum, current) => sum + current, 0);
  
  return totalSum - min - max;
}
