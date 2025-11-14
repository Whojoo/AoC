using System.Text.RegularExpressions;

using Shared;

namespace AoC2015;

public static partial class Day13
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day13.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  private readonly record struct HappinessRecord(string Target, int Happiness);

  public static int PartOne(string[] input)
  {
    Dictionary<string, List<HappinessRecord>> relationshipsDict = [];

    foreach (var line in input)
    {
      var match = HappinessRegex().Match(line);
      
      var rawValue = int.Parse(match.Groups["happinessAmount"].Value);
      var value = match.Groups["delta"].Value switch
      {
        "gain" => rawValue,
        "lose" => -rawValue,
        _ => throw new ArgumentException("unexpected delta")
      };

      var list = relationshipsDict.GetOrAdd(match.Groups["sourceName"].Value, []);
      list.Add(new HappinessRecord(match.Groups["targetName"].Value, value));
    }
    
    return GetMaxHappiness(relationshipsDict.First().Key, relationshipsDict, []);
  }

  public static int PartTwo(string[] input)
  {
    Dictionary<string, List<HappinessRecord>> relationshipsDict = [];
    var youSet = new HashSet<string>();

    foreach (var line in input)
    {
      var match = HappinessRegex().Match(line);
      
      var rawValue = int.Parse(match.Groups["happinessAmount"].Value);
      var value = match.Groups["delta"].Value switch
      {
        "gain" => rawValue,
        "lose" => -rawValue,
        _ => throw new ArgumentException("unexpected delta")
      };

      var list = relationshipsDict.GetOrAdd(match.Groups["sourceName"].Value, [new HappinessRecord("You", 0)]);
      youSet.Add(match.Groups["sourceName"].Value);
      list.Add(new HappinessRecord(match.Groups["targetName"].Value, value));
    }

    relationshipsDict["You"] = youSet.Select(x => new HappinessRecord(x, 0)).ToList();
    
    return GetMaxHappiness("You", relationshipsDict, []);
  }

  private static int GetMaxHappiness(
    string source,
    Dictionary<string, List<HappinessRecord>> relationshipsDict,
    List<string> seatedMembers)
  {
    var relationShips = relationshipsDict[source];
    List<string> newSeatedMembers = [..seatedMembers, source];
    
    if (relationShips.Count != seatedMembers.Count)
    {
      return relationShips
        .Where(x => !seatedMembers.Contains(x.Target))
        .Select(x => GetMaxHappiness(x.Target, relationshipsDict, newSeatedMembers))
        .Max();
    }

    return newSeatedMembers
      .Select((name, index) => (index, name))
      .Aggregate(0, (current, tuple) =>
      {
        (int index, string name) = tuple;
        var left = index + 1 == newSeatedMembers.Count ? 0 : index + 1;
        var right = index == 0 ? newSeatedMembers.Count - 1 : index - 1;

        var sourceList = relationshipsDict[name];
        var leftScore = sourceList.First(x => x.Target == newSeatedMembers[left]).Happiness;
        var rightScore = sourceList.First(x => x.Target == newSeatedMembers[right]).Happiness;
        return current + leftScore + rightScore;
      });
  }

  [GeneratedRegex(@"(?<sourceName>[a-zA-Z]+) would (?<delta>gain|lose) (?<happinessAmount>\d+) happiness units by sitting next to (?<targetName>[a-zA-Z]+).")]
  private static partial Regex HappinessRegex();
}