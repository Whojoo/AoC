using System.Resources;

namespace AoC2015;

public static class Day8
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day8.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    var totalInMemoryCharacters = input
      .Select(line =>
      {
        var quotes = CountOccurrences("\"", line);
        var escapedBackslash = CountOccurrences(@"\\", line);
        var escapedHex = CountOccurrences(@"\x", line);

        var cheatBackslash = CountOccurrences(@"\\\", line);
        var counterCheatBackslash = CountOccurrences(@"\\\\", line);
        
        var cheatHex = CountOccurrences(@"\\x", line);
        var counterCheatHex = CountOccurrences(@"\\\x", line);

        escapedBackslash -= cheatBackslash - counterCheatBackslash;
        escapedHex -= cheatHex - counterCheatHex;

        return line.Length - quotes - escapedBackslash - escapedHex * 3;
      })
      .Sum();

    var totalCharacters = input.Select(line => line.Length).Sum();

    return totalCharacters - totalInMemoryCharacters;
  }

  public static int PartTwo(string[] input) =>
    input
      .Select(line =>
      {
        const int defaultIncrease = 2;
        var quotes = CountOccurrences("\"", line);
        var backslashes = CountOccurrences("\\", line);

        return defaultIncrease + quotes + backslashes;
      })
      .Sum();
  

  private static int CountOccurrences(string search, string target)
  {
    var index = target.IndexOf(search, StringComparison.InvariantCulture);
    var count = 0;

    while (index >= 0)
    {
      count++;
      index = target.IndexOf(search, index + 1, StringComparison.InvariantCulture);
    }
    
    return count;
  }
}