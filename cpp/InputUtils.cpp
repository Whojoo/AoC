//
// Created by degie on 09/08/2025.
//

#include "InputUtils.h"

#include <fstream>

std::vector<std::string> read_input_file(const unsigned int year, const std::string &filename) {
    const auto fullFileName = "../" + std::to_string(year) + "/input/" + filename;
    std::ifstream inputFile (fullFileName);

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
