export function getRealFloor(n: number): number {
  if (n < 0) {
      return n; 
  } else if (n === 0) {
      return 0; 
  } else if (n < 13) {
      return n - 1; 
  } else {
      return n - 2; 
  }
}
