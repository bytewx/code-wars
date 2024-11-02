export function nextBigger(n: number): number {
  const digits: number[] = n.toString().split('').map(Number);
  const length: number = digits.length;
  
  let pivot: number = -1;
  
  for (let i = length - 2; i >= 0; i--) {
    if (digits[i] < digits[i + 1]) {
      pivot = i;
      break;
    }
  }
  
  if (pivot === -1) {
    return -1;
  }
  
  let successor = -1;
  
  for (let i = length - 1; i > pivot; i--) {
    if (digits[i] > digits[pivot]) {
      successor = i;
      break;
    }
  }
  
  [digits[pivot], digits[successor]] = [digits[successor], digits[pivot]];
  
  const newDigits: number[] = digits.slice(0, pivot + 1).concat(digits.slice(pivot + 1).reverse());

  const nextBiggestNumber: number = Number(newDigits.join(''));
  
  return nextBiggestNumber;
}
