export function points(games : string[]): number {
  return games.reduce((total, game) => {
    const [x, y] = game.split(':').map(Number);
    
    if (x > y) {
      return total += 3;
    } else if (x === y) {
      return total += 1;
    } else {
      return total;
    }
  }, 0);
}
