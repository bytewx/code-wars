export function isTriangle(a: number, b: number, c: number): boolean {
  if (a <= 0 || b <= 0 || c <= 0) {
    return false;
  }

  return a + b > c && a + c > b && b + c > a;
}
