export function switchItUp(intNumber: number): string {
  const numberWords: string[] = [
    "Zero",  
    "One",   
    "Two",   
    "Three", 
    "Four",  
    "Five", 
    "Six",   
    "Seven", 
    "Eight", 
    "Nine"  
  ];
  
  return numberWords[intNumber];
}
