using FluentAssertions;
using Microsoft.VisualStudio.TestPlatform.TestHost;

namespace AoC2024.Day5.Test;

public class UnitTest1
{
    [Fact]
    public void CalculatePart1_WhenExampleInput_ShouldGiveExampleResult()
    {
        List<string> input =
        [
            "47|53",
            "97|13",
            "97|61",
            "97|47",
            "75|29",
            "61|13",
            "75|53",
            "29|13",
            "97|29",
            "53|29",
            "61|53",
            "97|53",
            "61|29",
            "47|13",
            "75|47",
            "97|75",
            "47|61",
            "75|61",
            "47|29",
            "75|13",
            "53|13",
            "",
            "75,47,61,53,29",
            "97,61,53,29,13",
            "75,29,13",
            "75,97,47,61,53",
            "61,13,29",
            "97,13,75,29,47"
        ];

        var result = Handler.CalculatePart1(input);

        result.Should().Be(143);
    }

    [Fact]
    public void CalculatePart2_WhenExampleInput_ShouldGiveExampleResult()
    {
        List<string> input =
        [
            "47|53",
            "97|13",
            "97|61",
            "97|47",
            "75|29",
            "61|13",
            "75|53",
            "29|13",
            "97|29",
            "53|29",
            "61|53",
            "97|53",
            "61|29",
            "47|13",
            "75|47",
            "97|75",
            "47|61",
            "75|61",
            "47|29",
            "75|13",
            "53|13",
            "",
            "75,47,61,53,29",
            "97,61,53,29,13",
            "75,29,13",
            "75,97,47,61,53",
            "61,13,29",
            "97,13,75,29,47"
        ];

        var result = Handler.CalculatePart2(input);

        result.Should().Be(123);
    }
}