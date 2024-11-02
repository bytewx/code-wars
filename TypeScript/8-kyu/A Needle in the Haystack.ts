export function findNeedle(haystack: any[]): string {
  const index = haystack.indexOf("needle");
  return `found the needle at position ${index}`;
}
