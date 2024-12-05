function cakes(recipe, available) {
  let maxCakes = Infinity;

  for (let ingredient in recipe) {
    if (!available[ingredient]) {
      return 0;
    }

    const possibleCakes = Math.floor(available[ingredient] / recipe[ingredient]);
    maxCakes = Math.min(maxCakes, possibleCakes);
  }

  return maxCakes;
}
