export function betterThanAverage(classPoints: number[], yourPoints: number) : boolean {
  const totalPoints = classPoints.reduce((accum, score) => accum + score, 0);
  
  const average = (totalPoints + yourPoints) / (classPoints.length + 1);
  
  return yourPoints > average;
}
