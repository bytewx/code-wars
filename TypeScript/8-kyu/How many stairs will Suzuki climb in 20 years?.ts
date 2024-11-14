export function stairsIn20(stairs: number[][]): number {
    const oneYearTotal = stairs.flat().reduce((acc, num) => acc + num, 0);

    return oneYearTotal * 20;
}
