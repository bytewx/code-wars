export function toCamelCase(str: string): string {
  const words = str.split(/[-_]/);
  
  const result = words
    .map((word, index) => {
      if (index === 0) {
        return word;
      }
      
      return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();
    })
    .join('');

  return result;
}

export function isAlpha(char: string): boolean {
  return /^[a-zA-Z]$/.test(char);
}
