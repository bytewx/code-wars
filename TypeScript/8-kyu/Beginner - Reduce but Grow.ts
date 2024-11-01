export function grow(arr: number[]): number {
  return arr.reduce((mult, value) => mult * value, 1);
}
