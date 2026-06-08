#include <stdbool.h>
#include <stddef.h>

typedef bool (*Predicate)(int);

bool any(const int* arr, size_t size, Predicate fun)
{
  if (size == 0) {
    return false;
  }
  
  for (int i = 0; i < size; i++) {
    if (fun(arr[i])) {
      return true;
    }
  }
  
  return false;
}
