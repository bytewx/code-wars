package kata

func ZeroFuel(distanceToPump int, mpg int, fuelLeft int) bool {
    return mpg * fuelLeft >= distanceToPump
}
