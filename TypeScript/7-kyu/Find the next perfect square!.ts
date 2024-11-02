export function findNextSquare(sq: number): number {
  const root: number = Math.sqrt(sq);
  
  if (root % 1 !== 0) {
    return -1; 
  }
  
  const nextRoot: number = root + 1;
  return nextRoot * nextRoot;
}
