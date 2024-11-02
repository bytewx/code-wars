export function humanYearsCatYearsDogYears(humanYears: number): [number, number, number] {
  let catYears = 0;
  let dogYears = 0;
  
  if (humanYears === 1) {
    catYears = 15;
  } else if (humanYears === 2) {
    catYears = 15 + 9;
  } else {
    catYears = 15 + 9 + (humanYears - 2) * 4;
  }
  
  if (humanYears === 1) {
    dogYears = 15; 
  } else if (humanYears === 2) {
    dogYears = 15 + 9;
  } else {
    dogYears = 15 + 9 + (humanYears - 2) * 5;
  }

  return [humanYears, catYears, dogYears];
}
