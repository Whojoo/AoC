#include <iostream>

#include "../InputUtils.h"

class Present {
    int length;
    int width;
    int height;

public:
    Present(const int length, const int width, const int height) {
        this->length = length;
        this->width = width;
        this->height = height;
    }

    [[nodiscard]] int get_length() const {
        return length;
    }

    [[nodiscard]] int get_width() const {
        return width;
    }

    [[nodiscard]] int get_height() const {
        return height;
    }
};

std::vector<Present> parse_input(const std::vector<std::string> &input);

int part_1(const std::vector<std::string> &input) {
    const auto parsedInput = parse_input(input);

    int paperNeeded = 0;

    for (auto currentPresent : parsedInput) {
        int minSide = 0;
        int current = currentPresent.get_width() *  currentPresent.get_height();
        paperNeeded += current * 2;

        minSide = current;

        current = currentPresent.get_width() * currentPresent.get_length();
        paperNeeded += current * 2;
        minSide = minSide > current ? current : minSide;

        current = currentPresent.get_height() * currentPresent.get_length();
        paperNeeded += current * 2;
        minSide = minSide > current ? current : minSide;

        paperNeeded += minSide;
    }

    return paperNeeded;
}

int part_2(const std::vector<std::string> &input) {
    const auto parsedInput = parse_input(input);

    int ribbonLength = 0;
    int bowRibbon = 0;

    for (auto currentPresent : parsedInput) {
        int left = currentPresent.get_width() < currentPresent.get_height() ? currentPresent.get_width() : currentPresent.get_height();
        int right = currentPresent.get_height() < currentPresent.get_length() ? currentPresent.get_height() : currentPresent.get_length();

        if (left == right) {
            left = currentPresent.get_width() < currentPresent.get_length() ? currentPresent.get_width() : currentPresent.get_length();
        }

        ribbonLength += left + left + right + right;
        bowRibbon += currentPresent.get_height() * currentPresent.get_width() * currentPresent.get_length();
    }

    return ribbonLength + bowRibbon;
}

int main() {
    const auto input = read_input_file(2015, "day2.txt");

    std::cout << "Part 1: " << std::to_string(part_1(input)) << std::endl;
    std::cout << "Part 2: " << std::to_string(part_2(input)) << std::endl;

    return 0;
}

std::vector<Present> parse_input(const std::vector<std::string> &input) {
    std::vector<Present> parsedInput;
    parsedInput.reserve(input.size());

    for (const auto & i : input) {
        int length = 0;
        int width = 0;
        int height = 0;
        int j = 0;
        std::string currentNumber;

        for (; j < i.length(); j++) {
            if (i[j] == 'x')
                break;

            currentNumber.push_back(i[j]);
        }

        length = std::stoi(currentNumber);
        currentNumber = "";
        j++;

        for (; j < i.length(); j++) {
            if (i[j] == 'x')
                break;

            currentNumber.push_back(i[j]);
        }

        width = std::stoi(currentNumber);
        currentNumber = "";
        j++;

        for (; j < i.length(); j++) {
            if (i[j] == '\r')
                break;

            currentNumber.push_back(i[j]);
        }

        height = std::stoi(currentNumber);
        parsedInput.emplace_back(length, width, height);
    }

    return parsedInput;
}
