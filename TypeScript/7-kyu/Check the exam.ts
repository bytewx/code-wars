export function checkExam(array1: string[], array2: string[]): number {
    let score = 0;

    for (let i = 0; i < array1.length; i++) {
        if (array2[i] === "") {
            score += 0;
        } else if (array2[i] === array1[i]) {
            score += 4; 
        } else {
            score -= 1; 
        }
    }

    return score < 0 ? 0 : score;
}
