export class Kata {
  static squareDigits(num: number): number {
    const digits = num.toString().split('');

    const squaredDigits = digits.map(digit => Number(digit) * Number(digit)).join('');

    return parseInt(squaredDigits, 10);
  }
}
