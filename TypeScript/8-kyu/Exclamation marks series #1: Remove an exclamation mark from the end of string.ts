export function remove(s: string): string {
  if (s.endsWith('!')) {
    return s.slice(0, -1);
  }
  
  return s;
}
