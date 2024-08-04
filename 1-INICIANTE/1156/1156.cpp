#include <iostream>
#include <iomanip>

int main() {
    double s = 1.0;
    double x = 3.0;
    double y = 2.0;

    while (x <= 39.0) {
        s += x / y;
        x += 2.0;
        y *= 2.0;
    }

    std::cout << std::fixed << std::setprecision(2) << s << std::endl;

    return 0;
}