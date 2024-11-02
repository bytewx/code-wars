export const towerBuilder = (nFloors: number): string[] => {
  const tower: string[] = [];
  
  for (let i = 0; i < nFloors; i++) {
    const numStars = 2 * i + 1;
    const numSpaces = nFloors - i - 1;
    
    tower.push(" ".repeat(numSpaces) + "*".repeat(numStars) + " ".repeat(numSpaces));
  }
  
  return tower;
}
