function XO(str) {
    str = str.toLowerCase();

    let xCount = (str.match(/x/g) || []).length;
    let oCount = (str.match(/o/g) || []).length;

    return xCount === oCount;
}
