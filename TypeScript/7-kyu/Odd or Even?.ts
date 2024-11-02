export function oddOrEven(array: number[]) {
  const sum = array.reduce((acc, curr) => acc + curr, 0);
  
  return sum % 2 === 0 ? "even" : "odd";
}
