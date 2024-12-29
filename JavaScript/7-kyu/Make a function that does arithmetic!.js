function arithmetic(a, b, operator) {
  const operations = {
    add: (x, y) => x + y,
    subtract: (x, y) => x - y,
    multiply: (x, y) => x * y,
    divide: (x, y) => x / y,
  };

  return operations[operator](a, b);
}
