package kata

func SmallestIntegerFinder(numbers []int) int {
  smallest := numbers[0]

  for _, value := range numbers {
    if value < smallest {
      smallest = value
    }
  }

  return smallest
}
