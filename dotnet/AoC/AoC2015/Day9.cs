using System.Text.RegularExpressions;

namespace AoC2015;

public static partial class Day9
{
  private const string TownsFormat = "{0} {1}";
  
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day9.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    (Dictionary<string, int> distances, List<string> cities) = PopulateDistancesDict(input);
    var minimumDistance = int.MaxValue;

    foreach (var city in cities)
    {
      minimumDistance = Math.Min(minimumDistance, CalcPartOne(distances, [city], cities, city));
    }
    
    return minimumDistance;
  }

  private static int CalcPartOne(
    Dictionary<string, int> distances,
    List<string> usedCities,
    List<string> cities,
    string startCity,
    int currentDistance = 0,
    int maximumDistance = int.MaxValue)
  {
    if (cities.Count == usedCities.Count)
    {
      return currentDistance;
    }

    var newMax = maximumDistance;
    
    foreach (var city in cities.Except(usedCities))
    {
      var newDistance = currentDistance + distances[string.Format(TownsFormat, startCity, city)];
      if (newDistance > maximumDistance)
        continue;

      List<string> newUsedCities = [.. usedCities, city];
      var distance = CalcPartOne(distances, newUsedCities, cities, city, newDistance, newMax);
      newMax = Math.Min(distance, newMax);
    }

    return newMax;
  }
  
  public static int PartTwo(string[] input)
  {
    (Dictionary<string, int> distances, List<string> cities) = PopulateDistancesDict(input);
    var minimumDistance = int.MinValue;

    foreach (var city in cities)
    {
      minimumDistance = Math.Max(minimumDistance, CalcPartTwo(distances, [city], cities, city));
    }
    
    return minimumDistance;
  }

  private static int CalcPartTwo(
    Dictionary<string, int> distances,
    List<string> usedCities,
    List<string> cities,
    string startCity,
    int currentDistance = 0,
    int maximumDistance = int.MinValue)
  {
    if (cities.Count == usedCities.Count)
    {
      return currentDistance;
    }

    var newMax = maximumDistance;
    
    foreach (var city in cities.Except(usedCities))
    {
      var newDistance = currentDistance + distances[string.Format(TownsFormat, startCity, city)];
      if (newDistance < maximumDistance)
        continue;

      List<string> newUsedCities = [.. usedCities, city];
      var distance = CalcPartTwo(distances, newUsedCities, cities, city, newDistance, newMax);
      newMax = Math.Max(distance, newMax);
    }

    return newMax;
  }

  private static (Dictionary<string, int> distances, List<string> cities) PopulateDistancesDict(string[] input)
  {
    Dictionary<string, int> distances = [];
    HashSet<string> cities = [];

    foreach (var line in input)
    {
      var match = RetrieveCitiesRegex().Match(line);
      var distanceGroup = match.Groups["distance"];
      
      var distance = int.Parse(distanceGroup.Value);
      var left = match.Groups["left"].Value;
      var right = match.Groups["right"].Value;

      distances[string.Format(TownsFormat, left, right)] = distance;
      distances[string.Format(TownsFormat, right, left)] = distance;
      cities.Add(left);
      cities.Add(right);
    }

    return (distances, cities.ToList());
  }
  
  [GeneratedRegex(@"(?<left>\w+) to (?<right>\w+) = (?<distance>\d+)")]
  private static partial Regex RetrieveCitiesRegex();
}