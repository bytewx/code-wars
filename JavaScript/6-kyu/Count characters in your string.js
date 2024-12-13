function count(string) {
  const charCount = {};

  for (const char of string) {
    charCount[char] = (charCount[char] || 0) + 1;
  }

  return charCount;
}
