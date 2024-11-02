export function dutyFree(normPrice: number, discount: number, hol: number): number {
  const savingsPerBottle = (normPrice * discount) / 100; 
  return Math.floor(hol / savingsPerBottle);
}
