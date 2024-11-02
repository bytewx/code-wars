export function past(h: number, m: number, s: number): number {
  const hoursInMilliseconds = h * 3600000; 
  const minutesInMilliseconds = m * 60000; 
  const secondsInMilliseconds = s * 1000; 
  
  return hoursInMilliseconds + minutesInMilliseconds + secondsInMilliseconds;
}
