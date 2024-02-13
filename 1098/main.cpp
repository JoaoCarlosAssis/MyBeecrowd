#include <iostream>

int main(){
  float k = 0;
  float y = 0;
  float count = 0;
    for(int i = 0; i < 11; i++){
      for(int j = 1; j < 4; j++){
        std::cout << "I=" << k << " J=" << j + count << std::endl;
        y++;
      }
      k += 0.2;
      count += 0.2;
    }
  return 0;
}
