export function hello(name = ''): string {
  if (name === '') {
    return 'Hello, World!';
  }
  
  const capitalized = name.charAt(0).toUpperCase() + name.slice(1).toLowerCase();
  
  return `Hello, ${capitalized}!`;
}
