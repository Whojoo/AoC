using Shared;

namespace AoC2025;

public static class Day2
{
  public static void Run()
  {
    var lines = InputReader.ReadChallengeInput(2025, 2);

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static ulong PartOne(string[] input) =>
    input
      .MapToRanges()
      .RetrieveSplitPairs()
      .Where(IsIdMadeOfPairs)
      .Aggregate((ulong)0, (current, id) => current + id);

  public static ulong PartTwo(string[] input)
  {
    return input
      .MapToRanges()
      .RetrievePairs()
      .Aggregate((ulong)0, (current, next) => current + next);
  }
  
  private static IEnumerable<Range> MapToRanges(this string[] input)
  {
    return input
      .SelectMany(x => x.Split(','))
      .Where(x => !string.IsNullOrWhiteSpace(x))
      .Select(x => x.Split('-'))
      .Select(x => new Range(ulong.Parse(x[0]), ulong.Parse(x[1])));
  }

  extension(IEnumerable<Range> ranges)
  {
    private IEnumerable<ulong> RetrieveSplitPairs()
    {
      foreach (var range in ranges)
      {
        var current = range.From;
        while (current <= range.To)
        {
          yield return current;
          current++;
        }
      }
    }

    private IEnumerable<ulong> RetrievePairs()
    {
      foreach (var range in ranges)
      {
        var current = range.From;
        while (current <= range.To)
        {
          ulong digits = GetNumberOfDigits(current);
          
          if (IsIdMadeOfPairs(current, digits))
            yield return current;

          current++;
        }
      }
    }
  }
  
  private static bool IsIdMadeOfPairs(ulong id)
  {
    var digits = GetNumberOfDigits(id);
    var splitAt = digits / 2;

    return splitAt * 2 == digits && IsIdMadeOfPairs(id, digits, splitAt);
  }

  private static bool IsIdMadeOfPairs(ulong id, ulong digits)
  {
    ulong splitAt = digits / 2;

    while (splitAt > 0)
    {
      if (IsIdMadeOfPairs(id, digits, splitAt))
        return true;

      splitAt--;
    }

    return false;
  }

  private static bool IsIdMadeOfPairs(ulong id, ulong digits, ulong splitAt) => IsIdMadeOfPairs(id, digits, splitAt, []);

  private static bool IsIdMadeOfPairs(ulong id, ulong digits, ulong splitAt, ulong[] previousNumbers)
  {
    ulong rightDigits = digits;
    ulong remainingId = id;
    while (rightDigits >= splitAt)
    {
      var power = (ulong)Math.Pow(10, digits - splitAt);
      var left = remainingId / power;
      var right = remainingId - (left * power);

      rightDigits = digits - splitAt;

      if (rightDigits == splitAt) return left == right && previousNumbers.All(x => x == left);
      
      digits = rightDigits;
      previousNumbers = [left, ..previousNumbers];
      remainingId = right;
    }

    return false;
  }

  private static ulong GetNumberOfDigits(ulong number)
  {
    if (number <= 0) return 1;

    ulong digits = 1;
    ulong comparison = 10;

    while (comparison <= number)
    {
      digits++;
      comparison *= 10;
    }

    return digits;
  }

  readonly record struct Range(ulong From, ulong To);
}