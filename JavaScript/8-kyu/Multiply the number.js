function multiply(number) {
  const numDigits = Math.abs(number).toString().length;
  const multiplier = Math.pow(5, numDigits);
  return number * multiplier;
}
