export const distinct = (a: number[]): number[] => {
  const seen = new Set<number>();
  const result: number[] = [];

  for (const num of a) {
    if (!seen.has(num)) {
      seen.add(num); 
      result.push(num); 
    }
  }

  return result; 
};
