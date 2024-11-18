export function usdcny(usd: number): string {
  const conversionRate = 6.75;
  const cny = (usd * conversionRate).toFixed(2); 
  return `${cny} Chinese Yuan`;
}
