export function findUniq(arr: number[]): number {
  if (arr[0] !== arr[1] && arr[0] !== arr[2]) return arr[0];
  if (arr[1] !== arr[0] && arr[1] !== arr[2]) return arr[1];
  
  for (let i = 2; i < arr.length; i++) {
    if (arr[i] !== arr[0]) return arr[i];
  }
  
  return arr[0];
}
