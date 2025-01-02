using FluentAssertions;

namespace Dotnet.Test;

public class Day7Tests
{
    [Fact]
    public void Part1()
    {
        // Arrange
        List<string> rawInput =
        [
            "190: 10 19",
            "3267: 81 40 27",
            "83: 17 5",
            "156: 15 6",
            "7290: 6 8 6 15",
            "161011: 16 10 13",
            "192: 17 8 14",
            "21037: 9 7 18 13",
            "292: 11 6 16 20"
        ];

        var input = rawInput.Select(Day7.MapToEquation).ToList();

        // Act
        var result = Day7.CalculatePart1(input);

        // Assert
        result.Should().Be(3749);
    }
    
    [Fact]
    public void Part2()
    {
        // Arrange
        List<string> rawInput =
        [
            "190: 10 19",
            "3267: 81 40 27",
            "83: 17 5",
            "156: 15 6",
            "7290: 6 8 6 15",
            "161011: 16 10 13",
            "192: 17 8 14",
            "21037: 9 7 18 13",
            "292: 11 6 16 20"
        ];

        var input = rawInput.Select(Day7.MapToEquation).ToList();

        // Act
        var result = Day7.CalculatePart2(input);

        // Assert
        result.Should().Be(11387);
    }
}