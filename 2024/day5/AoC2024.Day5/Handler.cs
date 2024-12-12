using System.Runtime.CompilerServices;

namespace AoC2024.Day5;

public static class Handler
{
    public static int CalculatePart1(List<string> input)
    {
        var (rawRules, rawUpdates) = SplitInput(input);
        var rules = ParseRules(rawRules);
        var updates = ParseUpdates(rawUpdates);

        var sum = updates
            .Where(update => IsValidUpdate(update, rules))
            .Sum(update => update[update.Count / 2]);

        return sum;
    }

    public static int CalculatePart2(List<string> input)
    {
        var (rawRules, rawUpdates) = SplitInput(input);
        var rules = ParseRules(rawRules);
        var updates = ParseUpdates(rawUpdates);

        var sum = updates
            .Where(update => !IsValidUpdate(update, rules))
            .Select(update => update.OrderByRule(rules))
            .Sum(update => update[update.Count / 2]);

        return sum;
    }

    private static List<int> OrderByRule(this List<int> update, Dictionary<int, ISet<int>> rules) =>
        update
            .Order(Comparer<int>.Create((x, y) =>
            {
                if (rules.GetSetForKey(x).Contains(y))
                {
                    return -1;
                }

                return rules.GetSetForKey(y).Contains(x) ? 1 : 0;
            }))
            .ToList();

    private static bool IsValidUpdate(List<int> update, Dictionary<int, ISet<int>> rules)
    {
        var isValid = true;

        for (var i = update.Count - 1; i > 0; i--)
        {
            var updatePage = update[i];
            isValid = isValid && !update.Take(i).Any(rules.GetSetForKey(updatePage).Contains);
        }

        return isValid;
    }

    private static (List<string> rawRules, List<string> rawUpdates) SplitInput(List<string> input)
    {
        var splitIndex = input.IndexOf(string.Empty);

        return (input.Take(splitIndex).ToList(), input.Skip(splitIndex + 1).ToList());
    }

    private static Dictionary<int, ISet<int>> ParseRules(List<string> rawRules)
    {
        var rules = new Dictionary<int, ISet<int>>();

        foreach (var rule in rawRules)
        {
            var values = rule.Split("|").Select(int.Parse).ToArray();
            var set = rules.GetSetForKey(values[0]);
            set.Add(values[1]);
        }

        return rules;
    }

    private static List<List<int>> ParseUpdates(List<string> rawUpdates) =>
        rawUpdates
            .Select(update => update
                .Split(",")
                .Select(int.Parse)
                .ToList())
            .ToList();

    private static ISet<int> GetSetForKey(this Dictionary<int, ISet<int>> dictionary, int key)
    {
        if (dictionary.TryGetValue(key, out var value))
        {
            return value;
        }

        value = new HashSet<int>();
        dictionary.Add(key, value);

        return value;
    }
}