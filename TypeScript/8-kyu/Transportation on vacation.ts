export function rentalCarCost(d: number): number {
  const dailyRate = 40;
  let totalCost = d * dailyRate;

  if (d >= 7) {
    totalCost -= 50;
  } else if (d >= 3) {
    totalCost -= 20; 
  }

  return totalCost;
}
