function mergeArrays(arr1, arr2) {
  const mergedArray = [...new Set(arr1.concat(arr2))];
  
  mergedArray.sort((a, b) => a - b);
  
  return mergedArray;
}
