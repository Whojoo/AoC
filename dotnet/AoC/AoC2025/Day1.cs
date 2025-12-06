using Shared;

namespace AoC2025;

public static class Day1
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
    const char leftCharacter = 'L';

    foreach (var line in input)
    {
      var turningAmount = int.Parse(line[1..]);
      if (line[0] == leftCharacter) 
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
        const char leftCharacter = 'L';
        var turningAmount = int.Parse(x[1..]);
        var direction = x[0] == leftCharacter ? -1 : 1;
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
}