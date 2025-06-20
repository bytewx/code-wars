var encryptThis = function(text) {
  return text.split(' ').map(word => {
    if (word.length === 0) return '';
    
    const asciiCode = word.charCodeAt(0);
    
    if (word.length === 1) {
      return asciiCode.toString();
    } else if (word.length === 2) {
      return asciiCode.toString() + word[1];
    } else {
      return asciiCode.toString() + word[word.length - 1] + word.slice(2, -1) + word[1];
    }
  }).join(' ');
};
