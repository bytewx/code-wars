function sumDigPow(a, b) {
  let result = [];

  for (let num = a; num <= b; num++) {
    let digits = num.toString().split('');
    
    let sum = digits.reduce((acc, digit, index) => {
      return acc + Math.pow(Number(digit), index + 1);
    }, 0);

    if (sum === num) {
      result.push(num);
    }
  }

  return result;
}
