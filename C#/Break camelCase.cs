public class Kata
{
  public static string BreakCamelCase(string str)
  {
    string result = "";
    
    foreach (char character in str)
    {
      if (char.IsUpper(character))
      {
        result += " " + character;
      }
      else
      {
        result += character;
      }
    }
    
    return result;
  }
}
