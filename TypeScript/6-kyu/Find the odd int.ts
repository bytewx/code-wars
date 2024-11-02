export const findOdd = (xs: number[]): number => {
  return xs.reduce((accumulator, value) => accumulator ^ value, 0);
};
