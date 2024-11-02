export function order(words: string): string {
  if (words === '') {
    return '';
  }
  
  return words.split(' ').sort((a: string, b: string): number => {
    const matchA: RegExpMatchArray | null = a.match(/\d+/);
    const matchB: RegExpMatchArray | null = b.match(/\d+/);

    const numA: number = matchA ? parseInt(matchA[0]) : 0;
    const numB: number = matchB ? parseInt(matchB[0]) : 0;
    
    return numA - numB;
  }).join(' ');
}
