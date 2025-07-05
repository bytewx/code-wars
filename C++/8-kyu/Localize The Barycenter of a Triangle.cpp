#include <utility>
#include <cmath>
#include <iomanip>

using point = const std::pair<double, double>;

std::pair<double, double> barTriang(point p1, point p2, point p3) {
  double x = (p1.first + p2.first + p3.first) / 3;
  double y = (p1.second + p2.second + p3.second) / 3;

  x = std::round(x * 10000.0) / 10000.0;
  y = std::round(y * 10000.0) / 10000.0;

  return std::pair<double, double>(x, y);
}
