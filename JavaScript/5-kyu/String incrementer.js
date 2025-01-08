function incrementString(strng) {
  const match = strng.match(/(\d*)$/);
  const number = match[0];
  
  if (!number) {
    return strng + "1";
  }
  
  const numberLength = number.length; 
  const incrementedNumber = (parseInt(number, 10) + 1).toString();
  
  const paddedNumber = incrementedNumber.padStart(numberLength, "0");
  
  return strng.slice(0, -numberLength) + paddedNumber;
}
