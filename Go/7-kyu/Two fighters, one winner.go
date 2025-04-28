package kata

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
    var currentAttacker, currentDefender *Fighter

    if firstAttacker == fighter1.Name {
        currentAttacker = &fighter1
        currentDefender = &fighter2
    } else {
        currentAttacker = &fighter2
        currentDefender = &fighter1
    }
  
    for {
        currentDefender.Health -= currentAttacker.DamagePerAttack

        if currentDefender.Health <= 0 {
            return currentAttacker.Name
        }

        currentAttacker, currentDefender = currentDefender, currentAttacker
    }
}
