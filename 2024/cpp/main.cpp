#include <iostream>
#include <map>
#include <string>
#include <queue>

int main() {
    std::priority_queue<int> left, right;
    std::map<int, int> similarity_map;
    std::string current_line;

    while (std::getline(std::cin, current_line)) {
        const auto first_space = current_line.find_first_of(' ');
        const auto second_space = current_line.find_last_of(' ');

        int first_number = std::stoi(current_line.substr(0, first_space));
        int second_number = std::stoi(current_line.substr(second_space, current_line.length()));

        left.push(first_number);
        right.push(second_number);

        if (similarity_map[second_number] > 0) {
            similarity_map[second_number] += 1;
        } else {
            similarity_map.insert_or_assign(second_number, 1);
        }
    }

    int sum_distances = 0, similarity_score = 0;

    while (!left.empty()) {
        const int num1 = left.top();
        const int num2 = right.top();

        sum_distances += std::abs(num1 - num2);

        left.pop();
        right.pop();

        if (similarity_map[num1] > 0) {
            similarity_score += num1 * similarity_map[num1];
        }
    }

    std::cout << "Sum Distances: " << sum_distances << "\n";
    std::cout << "Similarity Score: " << similarity_score << std::endl;

    return 0;
}
