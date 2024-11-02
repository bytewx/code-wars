export const minValue = (values: number[]): number => {
  const uniqueValues = Array.from(new Set(values));

  uniqueValues.sort((a, b) => a - b);

  const smallestNumber = parseInt(uniqueValues.join(''), 10);

  return smallestNumber;
};
