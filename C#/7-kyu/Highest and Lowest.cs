using System;
using System.Linq;

public static class Kata
{
  public static string HighAndLow(string numbers)
  {
    int[] numArray = numbers.Split(' ').Select(int.Parse).ToArray();
        
    int maxNum = numArray.Max();
    int minNum = numArray.Min();

    return $"{maxNum} {minNum}";
  }
}
