using System;
using System.Linq;

public class Kata
{
  public static string DuplicateEncode(string word)
  {
    word = word.ToLower();

    return new string(word.Select(c => word.Count(x => x == c) > 1 ? ')' : '(').ToArray());
  }
}
