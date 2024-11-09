export function sumOfDifferences(arr: number[]): number {
  if (arr.length < 2) {
    return 0;
  }
  
  arr.sort((a, b) => b - a);
  
  let sum: number = 0;
  
  for (let i = 0; i < arr.length - 1; i++) {
    sum += arr[i] - arr[i + 1];
  }
  
  return sum;
}
