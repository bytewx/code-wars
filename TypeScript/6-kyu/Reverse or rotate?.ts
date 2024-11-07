export function revRot(s: string, sz: number): string {
  if (sz <= 0 || s === "" || sz > s.length) return "";

  let result = "";

  for (let i = 0; i < s.length; i += sz) {
    let chunk = s.slice(i, i + sz);

    if (chunk.length < sz) break;

    const sum = chunk.split('').reduce((acc, digit) => acc + parseInt(digit), 0);

    if (sum % 2 === 0) {
      result += chunk.split('').reverse().join('');
    } else {
      result += chunk.slice(1) + chunk[0];
    }
  }

  return result;
}
