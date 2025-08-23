using System.Text.RegularExpressions;

namespace AoC2015;

public static partial class Day6
{
    public static void Run()
    {
        var lines = File.ReadAllLines("input/day6.txt");
        
        Console.WriteLine($"Part one: {PartOne(lines)}");
        Console.WriteLine($"Part two: {PartTwo(lines)}");
    }
    
    public static int PartOne(string[] input)
    {
        var world = new bool[1000][];
        for (var i = 0; i < 1000; i++)
            world[i] = new bool[1000];

        var rules = ParseRules(input);
        
        foreach (var rule in rules)
            rule.PerformLightRule(world);

        return world
            .SelectMany(x => x)
            .Count(x => x);
    }
    
    public static int PartTwo(string[] input)
    {
        var world = new int[1000][];
        for (var i = 0; i < 1000; i++)
            world[i] = new int[1000];

        var rules = ParseBrightnessRules(input);
        
        foreach (var rule in rules)
            rule.AdjustBrightness(world);

        return world
            .SelectMany(x => x)
            .Sum();
    }

    private static ILightRule[] ParseRules(string[] input)
    {
        var rules = new ILightRule[input.Length];

        for (var i = 0; i < input.Length; i++)
        {
            var match = ParseInstructionsRegex().Match(input[i]);
            var start = new Position(int.Parse(match.Groups["StartX"].Value), int.Parse(match.Groups["StartY"].Value));
            var end = new Position(int.Parse(match.Groups["EndX"].Value), int.Parse(match.Groups["EndY"].Value));

            rules[i] = match.Groups["Rule"].Value switch
            {
                "turn on" => new ForceStateRule(start, end, true),
                "turn off" => new ForceStateRule(start, end, false),
                "toggle" => new ToggleStateRule(start, end),
                _ => rules[i]
            };
        }

        return rules;
    }
    
    private static IBrightnessRule[] ParseBrightnessRules(string[] input)
    {
        var rules = new IBrightnessRule[input.Length];

        for (var i = 0; i < input.Length; i++)
        {
            var match = ParseInstructionsRegex().Match(input[i]);
            var start = new Position(int.Parse(match.Groups["StartX"].Value), int.Parse(match.Groups["StartY"].Value));
            var end = new Position(int.Parse(match.Groups["EndX"].Value), int.Parse(match.Groups["EndY"].Value));

            rules[i] = match.Groups["Rule"].Value switch
            {
                "turn on" => new ForceBrightnessRule(start, end, 1),
                "turn off" => new ForceBrightnessRule(start, end, -1),
                "toggle" => new ToggleBrightnessRule(start, end),
                _ => rules[i]
            };
        }

        return rules;
    }

    private readonly record struct Position(int X, int Y);
    
    private interface ILightRule
    {
        void PerformLightRule(bool[][] world);
    }

    private class ToggleStateRule(Position startPos, Position endPos) : ILightRule
    {
        public void PerformLightRule(bool[][] world)
        {
            for (var y = startPos.Y; y <= endPos.Y; y++)
            for (var x = startPos.X; x <= endPos.X; x++)
                world[y][x] = !world[y][x];
        }
    }

    private class ForceStateRule(Position startPos, Position endPos, bool state) : ILightRule
    {
        public void PerformLightRule(bool[][] world)
        {
            for (var y = startPos.Y; y <= endPos.Y; y++)
            for (var x = startPos.X; x <= endPos.X; x++)
                world[y][x] = state;
        }
    }

    private interface IBrightnessRule
    {
        void AdjustBrightness(int[][] world);
    }

    private class ToggleBrightnessRule(Position startPos, Position endPos) : IBrightnessRule
    {
        public void AdjustBrightness(int[][] world)
        {
            for (var y = startPos.Y; y <= endPos.Y; y++)
            for (var x = startPos.X; x <= endPos.X; x++)
                world[y][x] += 2;
        }
    }

    private class ForceBrightnessRule(Position startPos, Position endPos, int adjustment) : IBrightnessRule
    {
        public void AdjustBrightness(int[][] world)
        {
            for (var y = startPos.Y; y <= endPos.Y; y++)
            for (var x = startPos.X; x <= endPos.X; x++)
                world[y][x] = Math.Max(0, world[y][x] + adjustment);
        }
    }

    [GeneratedRegex(@"(?<Rule>turn on|turn off|toggle) (?<StartX>[\d]{0,3}),(?<StartY>[\d]{0,3}) through (?<EndX>[\d]{0,3}),(?<EndY>[\d]{0,3})")]
    public static partial Regex ParseInstructionsRegex();
}