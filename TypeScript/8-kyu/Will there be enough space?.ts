export function enough(cap: number, on: number, wait: number): number {
  const totalPeople = on + wait; 
  const spaceLeft = cap - on;

  return totalPeople <= cap ? 0 : totalPeople - cap;
}
