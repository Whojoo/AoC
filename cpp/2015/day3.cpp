#include <fstream>
#include <iostream>
#include <memory>
#include <set>
#include <string>
#include <vector>

class Position {
    int x;
    int y;

public:
    Position(const int x, const int y)
        : x(x),
          y(y) {
    }

    [[nodiscard]] int get_x() const {
        return x;
    }

    [[nodiscard]] int get_y() const {
        return y;
    }

    friend bool operator<(const Position &lhs, const Position &rhs) {
        if (lhs.x < rhs.x)
            return true;
        if (rhs.x < lhs.x)
            return false;
        return lhs.y < rhs.y;
    }

    friend bool operator<=(const Position &lhs, const Position &rhs) {
        return rhs >= lhs;
    }

    friend bool operator>(const Position &lhs, const Position &rhs) {
        return rhs < lhs;
    }

    friend bool operator>=(const Position &lhs, const Position &rhs) {
        return !(lhs < rhs);
    }
};

int part_1(const std::vector<std::string> &input) {
    std::set<Position> visitedPositions;
    Position currentPosition(0, 0);
    visitedPositions.insert(currentPosition);

    for (auto const &line: input) {
        for (auto const &directionChar: line) {
            if (directionChar == '<')
                currentPosition = Position(currentPosition.get_x() - 1, currentPosition.get_y());
            else if (directionChar == '>')
                currentPosition = Position(currentPosition.get_x() + 1, currentPosition.get_y());
            else if (directionChar == '^')
                currentPosition = Position(currentPosition.get_x(), currentPosition.get_y() - 1);
            else if (directionChar == 'v')
                currentPosition = Position(currentPosition.get_x(), currentPosition.get_y() + 1);

            visitedPositions.insert(currentPosition);
        }
    }

    return visitedPositions.size();
}

int part_2(const std::vector<std::string> &input) {
    std::set<Position> visitedPositions;
    Position currentPosition(0, 0);
    Position robotPosition(0, 0);
    visitedPositions.insert(currentPosition);

    for (auto const &line: input) {
        for (int i = 0; i < line.length(); i += 2) {
            auto directionChar = line[i];
            if (directionChar == '<')
                currentPosition = Position(currentPosition.get_x() - 1, currentPosition.get_y());
            else if (directionChar == '>')
                currentPosition = Position(currentPosition.get_x() + 1, currentPosition.get_y());
            else if (directionChar == '^')
                currentPosition = Position(currentPosition.get_x(), currentPosition.get_y() - 1);
            else if (directionChar == 'v')
                currentPosition = Position(currentPosition.get_x(), currentPosition.get_y() + 1);

            directionChar = line[i + 1];
            if (directionChar == '<')
                robotPosition = Position(robotPosition.get_x() - 1, robotPosition.get_y());
            else if (directionChar == '>')
                robotPosition = Position(robotPosition.get_x() + 1, robotPosition.get_y());
            else if (directionChar == '^')
                robotPosition = Position(robotPosition.get_x(), robotPosition.get_y() - 1);
            else if (directionChar == 'v')
                robotPosition = Position(robotPosition.get_x(), robotPosition.get_y() + 1);

            visitedPositions.insert(currentPosition);
            visitedPositions.insert(robotPosition);
        }
    }

    return visitedPositions.size();
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
