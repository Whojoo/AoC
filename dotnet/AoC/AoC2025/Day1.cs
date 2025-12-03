using System.Text.RegularExpressions;

using Shared;

namespace AoC2025;

public static partial class Day1
{
  private const int StartingPosition = 50;
  private const int Top = 100;
  
  public static void Run()
  {
    var lines = InputReader.ReadChallengeInput(2025, 1);

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    var amountOfZeros = 0;
    var currentPosition = StartingPosition;

    foreach (var line in input)
    {
      var match = DirectionRegex().Match(line);
      var turningAmount = match.Groups["amount"].IntValue();
      if (match.Groups["direction"].Value == "L")
        turningAmount *= -1;

      currentPosition += turningAmount;

      if (currentPosition % Top == 0)
        amountOfZeros++;
    }

    return amountOfZeros;
  }

  public static int PartTwo(string[] input)
  {
    return input
      .SelectMany(x =>
      {
        var match = DirectionRegex().Match(x);
        var turningAmount = match.Groups["amount"].IntValue();
        var direction = match.Groups["direction"].Value == "L" ? -1 : 1;
        return Enumerable.Repeat(direction, turningAmount);
      })
      .Aggregate(new RotationInfo(StartingPosition, 0), (current, nextRotation) =>
      {
        var newCurrent = current with { CurrentPosition = current.CurrentPosition + nextRotation };

        if (newCurrent.CurrentPosition % Top == 0)
          newCurrent = newCurrent with { TotalZeros = newCurrent.TotalZeros + 1 };

        return newCurrent;
      })
      .TotalZeros;
  }

  private readonly record struct RotationInfo(int CurrentPosition, int TotalZeros);

  [GeneratedRegex(@"(?<direction>L|R)(?<amount>\d+)")]
  private static partial Regex DirectionRegex();
}