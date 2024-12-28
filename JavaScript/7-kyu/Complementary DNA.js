function dnaStrand(dna) {
  const complements = {
    'A': 'T',
    'T': 'A',
    'C': 'G',
    'G': 'C'
  };

  return dna.split('').map(base => complements[base]).join('');
}
