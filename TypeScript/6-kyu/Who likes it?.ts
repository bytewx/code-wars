export const likes = (a : string[]) : string => {
  const count: number = a.length;
  
  if (count === 0) {
    return "no one likes this";
  } else if (count === 1) {
    return `${a[0]} likes this`;
  } else if (count === 2) {
    return `${a[0]} and ${a[1]} like this`;
  } else if (count === 3) {
    return `${a[0]}, ${a[1]} and ${a[2]} like this`;
  } else {
    return `${a[0]}, ${a[1]} and ${count - 2} others like this`;
  }
}
