export function sumMix(x: any[]): number {
  return x.reduce((sum, value) => sum + parseInt(value), 0);
}
