function array(string) {
  const items = string.split(',');

  if (items.length < 3) {
    return null; 
  }

  return items.slice(1, -1).join(' ');
}
