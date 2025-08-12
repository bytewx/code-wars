#include <string>
#include <vector>
#include <algorithm>

std::string swap(const std::string& str)
{
  if (str.empty())
  {
    return "";  
  }
  
  std::string result = str;
  std::vector<char> vowels{ 'a', 'e', 'i', 'o', 'u' };
  
  for (char& character : result)
  {
    if (std::find(vowels.begin(), vowels.end(), character) != vowels.end())
    {
      character = static_cast<char>(std::toupper(static_cast<unsigned char>(character)));
    }
  }
  
  return result;
}
