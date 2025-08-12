#include <string>

std::string sum_str(const std::string& a, const std::string& b) {
  if (a.empty() && b.empty())
  {
    return "0";
  }
  
  if (a.empty())
  {
    return b;
  }
  
  if (b.empty())
  {
    return a;
  }
  
  return std::to_string(std::stoi(a) + std::stoi(b));
}
