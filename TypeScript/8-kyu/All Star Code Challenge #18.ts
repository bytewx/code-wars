export function strCount(str: string, letter: string): number {
  return str.split(letter).length - 1;
}
