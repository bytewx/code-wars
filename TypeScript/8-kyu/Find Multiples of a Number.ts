export function findMultiples(integer: number, limit: number): number[] {
  const multiples: number[] = [];
  
  for (let i = integer; i <= limit; i += integer) {
    multiples.push(i);
  }
  
  return multiples;
}
