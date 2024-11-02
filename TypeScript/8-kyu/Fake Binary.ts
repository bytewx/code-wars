export const fakeBin = (x:string):string => {
  return x.split('').map((digit) => (parseInt(digit) < 5 ? '0' : '1')).join('');
};
