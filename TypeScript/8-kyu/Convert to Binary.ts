export function toBinary(n: number): number {
  const binaryString = n.toString(2);
  
  return parseInt(binaryString, 10);
}
