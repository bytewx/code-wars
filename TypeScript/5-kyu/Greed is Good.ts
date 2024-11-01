const scoringRules: Record<string, number> = {
  "1-3": 1000,  
  "6-3": 600,   
  "5-3": 500,   
  "4-3": 400,   
  "3-3": 300,   
  "2-3": 200,   
  "1-1": 100,   
  "5-1": 50    
};

export function score(dice: number[]): number {
  const counts: Record<number, number> = {};

  for (const die of dice) {
    counts[die] = (counts[die] || 0) + 1;
  }

  let totalScore = 0;

  for (const key in scoringRules) {
    const [value, countNeeded] = key.split('-').map(Number);
    const points = scoringRules[key];

    if (counts[value] >= countNeeded) {
      totalScore += points;
      counts[value] -= countNeeded;
    }
  }
  
  totalScore += (counts[1] || 0) * 100; 
  totalScore += (counts[5] || 0) * 50;

  return totalScore;
}
