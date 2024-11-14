export function pipeFix(numbers: number[]): number[] {
  const start = numbers[0];
  const end = numbers[numbers.length - 1];
  const result = [];
  
  for (let i = start; i <= end; i++) {
    result.push(i);
  }
  
  return result;
}
