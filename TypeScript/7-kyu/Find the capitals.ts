export function capitals (word: string): number[] {
	const indices: number[] = [];

  for (let i = 0; i < word.length; i++) {
    if (word[i] === word[i].toUpperCase()) { 
      indices.push(i); 
    }
  }

  return indices;
}
