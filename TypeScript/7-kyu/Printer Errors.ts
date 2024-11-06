export function printerError(s: string): string {
  const errors = s.split('').filter(char => char < 'a' || char > 'm').length;
  
  return `${errors}/${s.length}`;
}
