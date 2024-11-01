export default function isSquare(n: number): boolean {
  if (n < 0) {
    return false;
  }
  
  return Number.isInteger(Math.sqrt(n));
};
