using System.Resources;

namespace AoC2015;

public static class Day8
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day8.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part one v2: {PartOneV2(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
    Console.WriteLine($"Part two v2: {PartTwoV2(lines)}");
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

  public static int PartOneV2(string[] input) =>
    input.Aggregate(0, (current, line) =>
    {
      var escapedCharacters = 0;
      for (var i = 0; i < line.Length; i++)
      {
        if (line[i] != '\\')
          continue;

        switch (line[i + 1])
        {
          case '\"':
          case '\\':
            i++;
            escapedCharacters++;
            break;
          case 'x':
            i += 3;
            escapedCharacters += 3;
            break;
        }
      }

      return current + escapedCharacters + 2;
    });

  public static int PartTwo(string[] input) =>
    input.Select(line =>
      {
        const int defaultIncrease = 2;
        var quotes = CountOccurrences("\"", line);
        var backslashes = CountOccurrences("\\", line);

        return defaultIncrease + quotes + backslashes;
      })
      .Sum();

  public static int PartTwoV2(string[] input) =>
    input.Aggregate(0, (current, line) =>
    {
      const int defaultIncrease = 2;
      var escapedCharacters = 0;
      foreach (var character in line)
        if (character is '\"' or '\\')
          escapedCharacters++;

      return defaultIncrease + escapedCharacters + current;
    });

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