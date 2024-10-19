using System;
using System.Collections.Generic;
using System.Linq;

public class Kata
{
  public static int DuplicateCount(string str)
  {
    string lowerInput = str.ToLower();

    var characterCounts = new Dictionary<char, int>();
    
    foreach (char character in lowerInput) {
      if (characterCounts.ContainsKey(character)) {
        characterCounts[character]++;
      } else {
        characterCounts[character] = 1;
      }
    }
    
    return characterCounts.Count(kvp => kvp.Value > 1);
  }
}
