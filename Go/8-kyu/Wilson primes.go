package kata

import "math/big"

func factorial(n int) *big.Int {
    if n <= 0 {
        return big.NewInt(1)
    }
  
    result := big.NewInt(1)
  
    for i := 2; i <= n; i++ {
        result.Mul(result, big.NewInt(int64(i)))
    }
  
    return result
}

func AmIWilson(n int) bool {
    if n <= 1 {
        return false
    }
    
    pMinus1Factorial := factorial(n - 1)
    
    pSquared := big.NewInt(int64(n * n))
    
    numerator := new(big.Int).Add(pMinus1Factorial, big.NewInt(1))
    
    return numerator.Mod(numerator, pSquared).Cmp(big.NewInt(0)) == 0
}
