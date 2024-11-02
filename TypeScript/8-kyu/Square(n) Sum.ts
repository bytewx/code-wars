export function squareSum(numbers: number[]): number {
  return numbers.reduce((accumulator, value) => accumulator + Math.pow(value, 2), 0);
}
