using System.Text.RegularExpressions;

namespace AoC2015;

public static partial class Day7
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day7.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    Dictionary<string, Node> nodesDictionary = ParseInput(input);
    Dictionary<string, ushort> cache = [];

    var startingNode = nodesDictionary["a"];

    return SolveForNode(startingNode, nodesDictionary, cache);
  }

  public static int PartTwo(string[] input)
  {
    Dictionary<string, Node> nodesDictionary = ParseInput(input);
    Dictionary<string, ushort> cache = [];

    var startingNode = nodesDictionary["a"];

    var result = SolveForNode(startingNode, nodesDictionary, cache);

    nodesDictionary["b"] = new Node("b") { LeftValue = result };
    cache.Clear();

    return SolveForNode(startingNode, nodesDictionary, cache);
  }

  private static ushort SolveForNode(Node targetNode, Dictionary<string, Node> nodesDictionary, Dictionary<string, ushort> cache)
  {
    if (cache.TryGetValue(targetNode.TargetName, out var cachedValue))
      return cachedValue;
    
    ushort? left = null;
    ushort? right = null;
    
    if (targetNode.HasLeftSide())
    {
      left = targetNode.LeftValue ?? SolveForNode(nodesDictionary[targetNode.LeftName!], nodesDictionary, cache);
    }

    if (targetNode.HasRightSide())
    {
      right = targetNode.RightValue ?? SolveForNode(nodesDictionary[targetNode.RightName!], nodesDictionary, cache);
    }

    ushort result = targetNode.Operator.HasValue ? ExecuteOperation(left, right, targetNode.Operator.Value) : (left ?? right)!.Value;
    cache[targetNode.TargetName] = result;
    return result;
  }

  private static ushort ExecuteOperation(ushort? left, ushort? right, LogicOperator logicOperator) =>
    logicOperator switch
    {
      LogicOperator.And => (ushort)(left!.Value & right!.Value),
      LogicOperator.Or => (ushort)(left!.Value | right!.Value),
      LogicOperator.Not => (ushort)(~ (left ?? right)!.Value),
      LogicOperator.RightShift => (ushort)(left!.Value >> right!.Value),
      LogicOperator.LeftShift => (ushort)(left!.Value << right!.Value),
      _ => throw new ArgumentOutOfRangeException(nameof(logicOperator), logicOperator, null)
    };

  private static Dictionary<string, Node> ParseInput(string[] input) =>
    input
      .Select(line =>
      {
        var match = PartOneRegex().Match(line);
        var node = new Node(match.Groups["target"].Value);

        if (match.Groups.TryGetValue("left", out var leftGroup) && !string.IsNullOrEmpty(leftGroup.Value))
        {
          if (ushort.TryParse(leftGroup.Value, out var leftValue))
            node.LeftValue = leftValue;
          else
            node.LeftName = leftGroup.Value;
        }

        if (match.Groups.TryGetValue("operator", out var operatorGroup) && !string.IsNullOrEmpty(operatorGroup.Value))
        {
          node.Operator = operatorGroup.Value switch
          {
            "AND" => LogicOperator.And,
            "OR" => LogicOperator.Or,
            "NOT" => LogicOperator.Not,
            "RSHIFT" => LogicOperator.RightShift,
            "LSHIFT" => LogicOperator.LeftShift,
            _ => throw new ArgumentOutOfRangeException()
          };
        }

        if (match.Groups.TryGetValue("right", out var rightGroup) && !string.IsNullOrEmpty(rightGroup.Value))
        {
          if (ushort.TryParse(rightGroup.Value, out var rightValue))
            node.RightValue = rightValue;
          else
            node.RightName = rightGroup.Value;
        }

        if (node.HasOperator() && node.Operator != LogicOperator.Not && !(node.HasLeftSide() && node.HasRightSide()))
        {
          throw new Exception("faaaaaaak");
        }

        return node;
      })
      .ToDictionary(node => node.TargetName);

  [GeneratedRegex(
    @"^(?<left>[a-z0-9]{0,5})[\s]?(?<operator>AND|RSHIFT|OR|LSHIFT|NOT)?[\s]?(?<right>[a-z0-9]{1,2})? -> (?<target>[a-z0-9]{1,2})$")]
  private static partial Regex PartOneRegex();

  private record Node(string TargetName)
  {
    public string? LeftName { get; set; }
    public ushort? LeftValue { get; set; }
    public LogicOperator? Operator { get; set; }
    public string? RightName { get; set; }
    public ushort? RightValue { get; set; }

    public bool HasLeftSide() => !string.IsNullOrEmpty(LeftName) || LeftValue.HasValue;
    public bool HasRightSide() => !string.IsNullOrEmpty(RightName) || RightValue.HasValue;
    public bool HasOperator() => Operator.HasValue;
  }

  private enum LogicOperator
  {
    And,
    Or,
    Not,
    RightShift,
    LeftShift,
  }
}