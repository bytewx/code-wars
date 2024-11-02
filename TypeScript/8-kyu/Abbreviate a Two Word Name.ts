export function abbrevName(name: string): string {
  const words = name.split(' ');
  
  const initials = `${words[0][0].toUpperCase()}.${words[1][0].toUpperCase()}`;

  return initials;
}
