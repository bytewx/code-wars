export function bump(x: string): string {
  const bumpCount = x.split('n').length - 1;
  return bumpCount <= 15 ? 'Woohoo!' : 'Car Dead';
}
