using System.Text.RegularExpressions;

namespace AoC2015;

public static partial class Day2
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day2.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input) => ParseInput(input)
    .Aggregate(0, (current, present) =>
    {
      var widthHeight = present.Width * present.Height;
      var widthLength = present.Width * present.Length;
      var heightLength = present.Height * present.Length;
      var minSide = Math.Min(Math.Min(widthHeight, widthLength), heightLength);
      return current + widthHeight * 2 + widthLength * 2 + heightLength * 2 + minSide;
    });
  

  public static int PartTwo(string[] input) => ParseInput(input)
    .Aggregate(0, (current, present) =>
    {
      var left = present.Width < present.Height ? present.Width : present.Height;
      var right = present.Height < present.Length ? present.Height : present.Length;

      if (left == right)
      {
        left = present.Width < present.Length ? present.Width : present.Length;
      }
      
      return current + left * 2 + right * 2 + present.Width * present.Height * present.Length;
    });


  private static Present[] ParseInput(string[] input) => input
    .Select(line =>
    {
      var match = InputRegex().Match(line);
      return new Present(
        match.Groups["length"].AsInteger(),
        match.Groups["width"].AsInteger(),
        match.Groups["height"].AsInteger());
    })
    .ToArray();

  private record Present(int Length, int Width, int Height);

  [GeneratedRegex(@"(?<length>\d+)x(?<width>\d+)x(?<height>\d+)")]
  private static partial Regex InputRegex();

  private static int AsInteger(this Group group) => int.Parse(group.Value);
}