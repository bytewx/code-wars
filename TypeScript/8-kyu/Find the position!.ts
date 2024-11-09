export function position(alphabet: string): string {
  let result: string = "";
  
  for (let i: number = 0; i < alphabet.length; i++) {
    let code: number = alphabet.toUpperCase().charCodeAt(i);
    
    if (code > 64 && code < 91) {
      result += (code - 64) + " ";
    } 
  }

  return `Position of alphabet: ${result.slice(0, result.length - 1)}`;
}
