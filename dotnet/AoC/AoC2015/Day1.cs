namespace AoC2015;

public static class Day1
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day1.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input) =>
    input[0].Aggregate(0, (current, floorMovement) => floorMovement switch
    {
      '(' => current + 1,
      ')' => current - 1,
      _ => throw new ArgumentOutOfRangeException(nameof(floorMovement), floorMovement, null)
    });

  public static int PartTwo(string[] input)
  {
    var currentFloor = 0;
    foreach ((int index, char floorMovement) in input[0].Select((character, i) => (i, character)))
    {
      currentFloor = floorMovement switch
      {
        '(' => currentFloor + 1,
        ')' => currentFloor - 1,
        _ => throw new ArgumentOutOfRangeException(nameof(floorMovement), floorMovement, null)
      };

      if (currentFloor == -1)
      {
        return index + 1;
      }
    }
    return 0;
  }
}