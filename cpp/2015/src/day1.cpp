#include <iostream>

#include "../../InputUtils.h"

int part_1(const std::vector<std::string> &input) {
    // Should be only a single line
    const auto& line = input[0];
    const auto length = line.length();

    int floor = 0;

    for (int i = 0; i < length; i++) {
        switch (const auto character = line[i]) {
            case '(': floor++; break;
            case ')': floor--; break;
            default: std::cout << "Uh oh at " << std::to_string(i) << " " << character << std::endl;
        }
    }

    return floor;
}

int part_2(const std::vector<std::string> &input) {
    // Should be only a single line
    const auto& line = input[0];
    const auto length = line.length();

    int floor = 0;

    for (int i = 0; i < length; i++) {
        switch (const auto character = line[i]) {
            case '(': floor++; break;
            case ')': floor--; break;
            default: std::cout << "Uh oh at " << std::to_string(i) << " " << character << std::endl;
        }

        if (floor == -1)
            return i+1;
    }

    return floor;
}

int main () {
    const auto input = read_input_file(2015, "day1.txt");

    std::cout << "Part 1: " << std::to_string(part_1(input)) << std::endl;
    std::cout << "Part 2: " << std::to_string(part_2(input)) << std::endl;
}
