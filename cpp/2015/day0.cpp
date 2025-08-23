#include <fstream>
#include <iostream>
#include <string>
#include <vector>

int part_1(const std::vector<std::string>& input) {
    return 0;
}

int part_2(const std::vector<std::string>& input) {
    return 0;
}

std::vector<std::string> read_input_file(const unsigned int year, const std::string &filename) {
    const auto fullFileName = std::to_string(year) + "/input/" + filename;
    std::ifstream inputFile(fullFileName);

    std::string line;
    std::vector<std::string> lines;
    while (std::getline(inputFile, line)) {
        // UTF-8 stuff, ignore the first 3 characters
        if (line[0] == '\357') {
            line = line.substr(3);
        }
        lines.push_back(line);
    }
    return lines;
}

int main() {
    const auto input = read_input_file(2015, "day3.txt");

    std::cout << "Part 1: " << std::to_string(part_1(input)) << std::endl;
    std::cout << "Part 2: " << std::to_string(part_2(input)) << std::endl;

    return 0;
}
