function sumPairs(ints, s) {
  const seen = new Set();
  
  for (const num of ints) {
    const complement = s - num;
    if (seen.has(complement)) {
      return [complement, num];
    }
    seen.add(num);
  }
  
  return undefined;
}
