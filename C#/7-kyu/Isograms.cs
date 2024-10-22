using System;
using System.Collections.Generic;

public class Kata
{
  public static bool IsIsogram(string str) 
  {
    str = str.ToLower();
    
    HashSet<char> seenLetters = new HashSet<char>();
    
    foreach (char c in str) 
    {
      if (seenLetters.Contains(c))
      {
        return false;  
      }
      
      seenLetters.Add(c);
    }
    
    return true;
  }
}
