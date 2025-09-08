function permutations(string) {
  if (string.length <= 1) return [string];

  const result = new Set();

  for (let i = 0; i < string.length; i++) {
    const char = string[i];

    const rest = string.slice(0, i) + string.slice(i + 1);

    for (let perm of permutations(rest)) {
      result.add(char + perm);
    }
  }

  return [...result];
}
