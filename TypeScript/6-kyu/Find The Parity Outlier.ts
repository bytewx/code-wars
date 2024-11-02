export function findOutlier(integers: number[]): number {
  let oddCount: number = 0;
  let evenCount: number = 0;
  let lastOdd: number | null = null;
  let lastEven: number | null = null;
  
  for (const num of integers) {
    if (num % 2 === 0) { 
        evenCount++;
        lastEven = num;
    } else { 
        oddCount++;
        lastOdd = num;
    }
  }

    return evenCount > oddCount ? lastOdd! : lastEven!;
}
