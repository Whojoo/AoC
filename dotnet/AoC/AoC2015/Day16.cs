using System.Text.RegularExpressions;

using Shared;

namespace AoC2015;

public static partial class Day16
{
  private const string IndexKey = "Index";
  private const string CatsKey = "cats";
  private const string TreesKey = "trees";
  private const string PomeraniansKey = "pomeranians";
  private const string GoldfishKey = "goldfish";

  public static void Run()
  {
    var lines = InputReader.ReadChallengeInput(2015, 16);

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    var originalSue = new Dictionary<string, int>
    {
      { "children", 3 },
      { "cats", 7 },
      { "samoyeds", 2 },
      { "pomeranians", 3 },
      { "akitas", 0 },
      { "vizslas", 0 },
      { "goldfish", 5 },
      { "trees", 3 },
      { "cars", 2 },
      { "perfumes", 1 }
    };

    return input
      .Select(ParseSue)
      .First(originalSue.CorrespondsWith)
      .GetValueOrDefault(IndexKey);
  }

  public static int PartTwo(string[] input)
  {
    var originalSue = new Dictionary<string, int>
    {
      { "children", 3 },
      { "cats", 7 },
      { "samoyeds", 2 },
      { "pomeranians", 3 },
      { "akitas", 0 },
      { "vizslas", 0 },
      { "goldfish", 5 },
      { "trees", 3 },
      { "cars", 2 },
      { "perfumes", 1 }
    };

    return input
      .Select(ParseSue)
      .First(originalSue.CorrespondsWithPlusPlus)
      .GetValueOrDefault(IndexKey);
  }

  private static Dictionary<string, int> ParseSue(string input)
  {
    // 0 = entire match and named groups are at the bottom, so our properties are on 1-6
    var match = GetSueRegex().Match(input);
    
    var sue = new Dictionary<string, int>
    {
      [IndexKey] = match.Groups["SueIndex"].IntValue(),
      [match.Groups[1].Value] = match.Groups[2].IntValue(),
      [match.Groups[3].Value] = match.Groups[4].IntValue(),
      [match.Groups[5].Value] = match.Groups[6].IntValue(),
    };

    return sue;
  }

  [GeneratedRegex(
    @"Sue (?<SueIndex>\d{1,3}): (?<FirstProperty>(\w+): (\d+),?\s?)(?<SecondProperty>(\w+): (\d+),?\s?)(?<ThirdProperty>(\w+): (\d+),?\s?)")]
  private static partial Regex GetSueRegex();

  extension(Dictionary<string, int> originalSue)
  {
    private bool CorrespondsWith(Dictionary<string, int> otherSue) =>
      originalSue.All(kp => !otherSue.TryGetValue(kp.Key, out var otherSueValue) || kp.Value == otherSueValue);

    private bool CorrespondsWithPlusPlus(Dictionary<string, int> otherSue)
    {
      var comparisonSue = originalSue.ToDictionary();
      
      if (otherSue.TryGetValue(CatsKey, out var otherSueCats) && comparisonSue[CatsKey] >= otherSueCats)
        return false;

      if (otherSue.TryGetValue(TreesKey, out var otherSueTrees) && comparisonSue[TreesKey] >= otherSueTrees)
        return false;

      if (otherSue.TryGetValue(PomeraniansKey, out var otherSuePom) && comparisonSue[PomeraniansKey] <= otherSuePom)
        return false;

      if (otherSue.TryGetValue(GoldfishKey, out var otherSueGoldfish) && comparisonSue[GoldfishKey] <= otherSueGoldfish)
        return false;

      comparisonSue.Remove(CatsKey);
      comparisonSue.Remove(TreesKey);
      comparisonSue.Remove(PomeraniansKey);
      comparisonSue.Remove(GoldfishKey);
      
      return comparisonSue.CorrespondsWith(otherSue);
    }
  }
}