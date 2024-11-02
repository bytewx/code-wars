export const high = (str: string): string => {
  const words = str.split(' ');

  let highestScore = 0;
  let highestScoringWord = '';

  for (const word of words) {
    const score = [...word].reduce((total, char) => total + (char.charCodeAt(0) - 96), 0);

    if (score > highestScore) {
      highestScore = score;
      highestScoringWord = word;
    }
  }

  return highestScoringWord;
};
