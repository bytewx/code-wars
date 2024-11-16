export function vowelIndices(word: string): number[] {
  const vowels = "aeiouyAEIOUY";
  
  return word
    .split("")
    .map((char, index) => vowels.includes(char) ? index + 1 : -1) 
    .filter(index => index !== -1); 
}
