export const longest = (s1:string, s2:string): string => {
  const combined: string = s1 + s2;
  
  const uniqueChars: Set<string> = new Set(combined);
  
  return Array.from(uniqueChars).sort().join('');
}
