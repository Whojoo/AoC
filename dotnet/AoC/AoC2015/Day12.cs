using System.Text.Json;
using System.Text.Json.Nodes;
using System.Text.RegularExpressions;

namespace AoC2015;

public static partial class Day12
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day12.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input) =>
    NumbersRegex()
      .Matches(input[0])
      .Aggregate(0, (current, match) => current + int.Parse(match.Value));

  public static int PartTwo(string[] input)
  {
    const string forbiddenValue = "red";
    var json = JsonNode.Parse(input[0])!;

    var nodeStack = new Stack<JsonNode>();
    nodeStack.Push(json);

    var total = 0;

    while (nodeStack.TryPop(out var node))
    {
      switch (node)
      {
        case JsonValue nodeValue when nodeValue.TryGetValue(out int numberValue):
          total += numberValue;
          break;
        case JsonObject nodeObject:
          {
            var hasRed = nodeObject
              .Select(x => x.Value)
              .Where(x => x is JsonValue value && value.GetValueKind() == JsonValueKind.String)
              .Any(x => string.Equals(forbiddenValue, x!.AsValue().ToString()));

            if (hasRed)
              break;
        
            foreach (var childNode in nodeObject.Where(x => x.Value is not null))
            {
              nodeStack.Push(childNode.Value!);
            }

            break;
          }
        case JsonArray nodeArray:
          {
            foreach (var childNode in nodeArray.Where(x => x is not null))
            {
              nodeStack.Push(childNode!);
            }

            break;
          }
      }
    }
    
    return total;
  }

  [GeneratedRegex(@"(\d+|-\d+)")]
  private static partial Regex NumbersRegex();
}