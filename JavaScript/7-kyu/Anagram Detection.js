var isAnagram = function(test, original) {
    const sortedTest = test.toLowerCase().split('').sort().join('');
    const sortedOriginal = original.toLowerCase().split('').sort().join('');
    
    return sortedTest === sortedOriginal;
};
