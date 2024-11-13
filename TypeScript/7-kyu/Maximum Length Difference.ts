export const mxdiflg = (a1: string[], a2: string[]): number => {
  if (a1.length === 0 || a2.length === 0) return -1;

  const maxLen1 = Math.max(...a1.map(str => str.length));
  const minLen1 = Math.min(...a1.map(str => str.length));
  const maxLen2 = Math.max(...a2.map(str => str.length));
  const minLen2 = Math.min(...a2.map(str => str.length));

  return Math.max(Math.abs(maxLen1 - minLen2), Math.abs(maxLen2 - minLen1));
};
