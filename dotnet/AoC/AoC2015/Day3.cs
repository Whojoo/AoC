namespace AoC2015;

public static class Day3
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day3.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    var currentPosition = new Position(0, 0);
    HashSet<Position> visitedPositions = [currentPosition];

    foreach (var direction in input[0])
    {
      currentPosition = Move(currentPosition, direction);
      visitedPositions.Add(currentPosition);
    }
    
    return visitedPositions.Count;
  }

  public static int PartTwo(string[] input)
  {
    var santaPosition = new Position(0, 0);
    var robotPosition = new Position(0, 0);
    HashSet<Position> visitedPositions = [santaPosition];

    foreach ((int index, char direction) in input[0].Select((dir, i) => (i, dir)))
    {
      if (index % 2 == 0)
      {
        santaPosition = Move(santaPosition, direction);
        visitedPositions.Add(santaPosition);
      }
      else
      {
        robotPosition = Move(robotPosition, direction);
        visitedPositions.Add(robotPosition);
      }
    }

    return visitedPositions.Count;
  }

  private static Position Move(Position currentPosition, char direction) =>
    direction switch
    {
      '<' => currentPosition with { X = currentPosition.X - 1 },
      '>' => currentPosition with { X = currentPosition.X + 1 },
      '^' => currentPosition with { Y = currentPosition.Y - 1 },
      'v' => currentPosition with { Y = currentPosition.Y + 1 },
      _ => throw new ArgumentOutOfRangeException()
    };

  private record Position(int X, int Y);
}