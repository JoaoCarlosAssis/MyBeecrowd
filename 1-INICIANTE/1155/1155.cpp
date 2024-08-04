#include <iostream>
#include <iomanip>

int main(){
  double s = 0.0;

  for(int i = 1; i<=100; i++){
    s += 1.0/i;
  }

  std::cout << std::fixed << std::setprecision(2) << s << std::endl;

  return 0;
}