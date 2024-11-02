export function findNb(m: number): number {
  let totalVolume = 0;
  let n = 0;

  while (totalVolume < m) {
    n++;
    totalVolume += n ** 3; 
  }

  return totalVolume === m ? n : -1;
}
