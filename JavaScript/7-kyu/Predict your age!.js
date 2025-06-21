function predictAge(...ages) {
  const sumOfSquares = ages.reduce((sum, age) => sum + age ** 2, 0);
  return Math.floor(Math.sqrt(sumOfSquares) / 2);
}
