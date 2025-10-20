using System.Text;

namespace AoC2015;

public static class Day10
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day10.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    return PerformLookAndSay(input[0], 40);
  }

  public static int PartTwo(string[] input)
  {
    return PerformLookAndSay(input[0], 50);
  }

  private static int PerformLookAndSay(string input, int iterations)
  {
    var line = new string(input);

    for (int i = 0; i < iterations; i++)
    {
      var newLine = new StringBuilder();

      for (int j = 0; j < line.Length;)
      {
        var character = line[j];
        var counter = 1;

        while (j + counter < line.Length && line[j + counter] == character)
        {
          counter++;
        }

        newLine.Append($"{counter}{character}");
        j += counter;
      }

      line = newLine.ToString();
    }
    
    return line.Length;      
  }
}