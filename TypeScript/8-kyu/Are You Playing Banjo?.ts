export function areYouPlayingBanjo(name: string): string {
  return name.toLowerCase().startsWith('r') ? `${name} plays banjo` : `${name} does not play banjo`;
}
