function comp(array1, array2) {
  if (!array1 || !array2) return false;
  if (array1.length !== array2.length) return false;
  
  let sortedArray1 = array1.map(x => x * x).sort((a, b) => a - b);
  let sortedArray2 = array2.sort((a, b) => a - b);
  
  return sortedArray1.every((val, index) => val === sortedArray2[index]);
}
