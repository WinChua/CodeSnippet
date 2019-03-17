#include <algorithm>
#include <random>

#include <string>
#include <vector>

int main() {
    std::vector<std::string> vec;
    std::mt19937_64 gen(43);

    for (int count = 0; count < 3; ++count) {
        for (int i=0; i < 10 * 1000 * 1000; ++i) {
            char buf[64];
            snprintf(buf, sizeof buf, "%016lx", gen());
            vec.push_back(buf);
        }
        std::sort(vec.begin(), vec.end());
        vec.clear();
    }
}
