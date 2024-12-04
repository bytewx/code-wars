function rot13(str) {
  return str.replace(/[A-Za-z]/g, char => {
    const base = char <= 'Z' ? 'A'.charCodeAt(0) : 'a'.charCodeAt(0);
    return String.fromCharCode(((char.charCodeAt(0) - base + 13) % 26) + base);
  });
}
