export function countPositivesSumNegatives(input: any): number[] {
  if (!input || input.length === 0) return [];

  let countPositives = 0;
  let sumNegatives = 0;

  for (let num of input) {
    if (num > 0) countPositives++;
    else if (num < 0) sumNegatives += num;
  }

  return [countPositives, sumNegatives];
}
