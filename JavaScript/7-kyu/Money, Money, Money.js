function calculateYears(principal, interest, tax, desired) {
    let years = 0;
    
    while (principal < desired) {
        let yearlyInterest = principal * interest;       
        let taxedInterest = yearlyInterest * (1 - tax);  
        principal += taxedInterest;                      
        years++;                                         
    }
    
    return years;
}
