export function isValidWalk(walk: string[]): boolean {
  if (walk.length !== 10) {
    return false;
  }
  
  let northSouth: number = 0;
  let eastWest: number = 0;
  
  for (const direction of walk) {
    switch (direction) {
      case 'n':
        northSouth += 1; 
        break;
      case 's':
        northSouth -= 1; 
        break;
      case 'e':
        eastWest += 1; 
        break;
      case 'w':
        eastWest -= 1; 
        break;
    }
  }
  
  return northSouth === 0 && eastWest === 0;
}
