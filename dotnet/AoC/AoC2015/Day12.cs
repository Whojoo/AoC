using System.Text.Json;
using System.Text.Json.Nodes;
using System.Text.RegularExpressions;

namespace AoC2015;

public static partial class Day12
{
  private const string ForbiddenValue = "red";
  
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day12.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part one: {PartOneFunctional(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
    Console.WriteLine($"Part two: {PartTwoFunctional(lines)}");
  }

  public static int PartOne(string[] input) =>
    NumbersRegex()
      .Matches(input[0])
      .Aggregate(0, (current, match) => current + int.Parse(match.Value));

  public static int PartTwo(string[] input)
  {
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
              .Any(x => string.Equals(ForbiddenValue, x!.AsValue().ToString()));

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

  public static int PartOneFunctional(string[] input) => ParseTotalFromNode(JsonNode.Parse(input[0]));
  
  public static int PartTwoFunctional(string[] input) => ParseTotalFromNodeWithoutForbiddenValue(JsonNode.Parse(input[0]));
  
  private static int ParseTotalFromNode(JsonNode? node) => node switch
  {
    JsonValue nodeValue => ParseTotalFromValue(nodeValue),
    JsonArray nodeArray => nodeArray.Aggregate(0, (current, childNode) => current + ParseTotalFromNode(childNode)),
    JsonObject nodeObject => nodeObject.Aggregate(0, (current, childNode) => current + ParseTotalFromNode(childNode.Value)),
    _ => 0
  };

  private static int ParseTotalFromNodeWithoutForbiddenValue(JsonNode? node) => node switch
  {
    JsonValue nodeValue => ParseTotalFromValue(nodeValue),
    JsonArray nodeArray => nodeArray.Aggregate(0, (current, childNode) => current + ParseTotalFromNodeWithoutForbiddenValue(childNode)),
    JsonObject nodeObject => ParseTotalFromObjectWithoutForbiddenValue(nodeObject),
    _ => 0
  };

  private static int ParseTotalFromValue(JsonValue nodeValue) =>
    nodeValue.GetValueKind() == JsonValueKind.Number ? nodeValue.GetValue<int>() : 0;

  private static int ParseTotalFromObjectWithoutForbiddenValue(JsonObject nodeObject) =>
    HasForbiddenValue(nodeObject) 
      ? 0 
      : nodeObject.Aggregate(0, (current, childNode) => current + ParseTotalFromNodeWithoutForbiddenValue(childNode.Value));

  private static bool HasForbiddenValue(JsonObject nodeObject) => 
    nodeObject
      .Select(x => x.Value)
      .Where(x => x is JsonValue value && value.GetValueKind() == JsonValueKind.String)
      .Any(x => string.Equals(ForbiddenValue, x!.AsValue().ToString()));

  [GeneratedRegex(@"(\d+|-\d+)")]
  private static partial Regex NumbersRegex();
}