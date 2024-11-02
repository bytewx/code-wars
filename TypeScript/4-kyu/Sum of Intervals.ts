export function sumOfIntervals(intervals: [number, number][]): number {
  if (intervals.length === 0) return 0;

  intervals.sort((a, b) => a[0] - b[0]);

  let totalLength: number = 0;
  let currentStart: number = intervals[0][0];
  let currentEnd: number = intervals[0][1];   

  for (let i = 1; i < intervals.length; i++) {
    const [start, end] = intervals[i];

    if (start > currentEnd) {
      totalLength += currentEnd - currentStart;
      currentStart = start;
      currentEnd = end;
    } else {
      currentEnd = Math.max(currentEnd, end);
    }
  }

  totalLength += currentEnd - currentStart;

  return totalLength;
}
