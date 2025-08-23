namespace AoC2015;

public static class Day5
{
    public static void Run()
    {
        var lines = File.ReadAllLines("input/day5.txt");

        Console.WriteLine($"Part one: {PartOne(lines)}");
        // Console.WriteLine($"Part two test: {PartTwo(["qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"])}");
        Console.WriteLine($"Part two: {PartTwo(lines)}");
    }

    public static int PartOne(string[] input)
    {
        IRule[] rules =
        [
            new ForbiddenStringsRule(["ab", "cd", "pq", "xy"]),
            new PairsRule(),
            new CountRequiredCharactersRule("aoiue", 3)
        ];

        return input.Count(x => rules.All(rule => rule.IsSatisfied(x)));
    }

    public static int PartTwo(string[] input)
    {
        IRule[] rules =
        [
            new DoubleNonOverlappingAnyPairs(),
            new SurroundedByPairRule()
        ];

        return input.Count(x => rules.All(rule => rule.IsSatisfied(x)));
    }

    private interface IRule
    {
        bool IsSatisfied(string input);
    }

    private class DoubleNonOverlappingAnyPairs : IRule
    {
        public bool IsSatisfied(string input)
        {
            HashSet<string> processedStrings = [];
            
            for (var i = 0; i < input.Length - 1; i++)
            {
                if (!processedStrings.Add(input[new Range(i, i + 2)]))
                    return true;

                if (input[i] == input[i + 1] && i + 2 < input.Length && input[i + 2] == input[i])
                    i++;
            }

            return false;
        }
    }

    private class SurroundedByPairRule : IRule
    {
        public bool IsSatisfied(string input)
        {
            for (var i = 1; i < input.Length - 1; i++)
            {
                if (input[i - 1] == input[i + 1])
                    return true;
            }

            return false;
        }
    }

    private class CountRequiredCharactersRule(string characterCollection, int minimalRequired) : IRule
    {
        public bool IsSatisfied(string input) => input.Count(characterCollection.Contains) >= minimalRequired;
    }

    private class PairsRule : IRule
    {
        public bool IsSatisfied(string input)
        {
            for (var i = 0; i < input.Length - 1; i++)
            {
                if (input[i] == input[i + 1])
                    return true;
            }

            return false;
        }
    }

    private class ForbiddenStringsRule(string[] forbiddenCharacters) : IRule
    {
        public bool IsSatisfied(string input) => forbiddenCharacters.All(x => !input.Contains(x));
    }
}