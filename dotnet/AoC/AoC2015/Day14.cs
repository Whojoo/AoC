using System.Text.RegularExpressions;

namespace AoC2015;

public static partial class Day14
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day14.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input) => ReindeerRace(input, 2503);

  public static int PartTwo(string[] input) => AdvancedReindeerRace(input, 2503);

  public static int ReindeerRace(string[] input, int totalTime) =>
    input
      .Select(x =>
      {
        var match = ReindeerStatsRegex().Match(x);
        var runningTime = int.Parse(match.Groups["runningTime"].Value);
        var restTime = int.Parse(match.Groups["restTime"].Value);
        var totalRunTime = runningTime + restTime;

        var remainingTime = totalTime % totalRunTime;
        var completeRuns = totalTime / totalRunTime;

        var speed = int.Parse(match.Groups["speed"].Value);
        var distanceInCompleteRuns = completeRuns * speed * runningTime;
        var distanceInRemainingTime = Math.Min(remainingTime, runningTime) * speed;

        return distanceInRemainingTime + distanceInCompleteRuns;
      })
      .Max();

  public static int AdvancedReindeerRace(string[] input, int totalTime)
  {
    var reindeer = input
      .Select(x =>
      {
        var match = ReindeerStatsRegex().Match(x);
        var runningTime = int.Parse(match.Groups["runningTime"].Value);
        var restTime = int.Parse(match.Groups["restTime"].Value);

        var speed = int.Parse(match.Groups["speed"].Value);

        return new ReindeerStats(runningTime, restTime, speed);
      })
      .ToArray();

    foreach (var second in Enumerable.Range(1, totalTime))
    {
      foreach (var singleReindeer in reindeer)
        singleReindeer.ProcessDistanceForSecond(second);

      var topDistance = reindeer.Max(x => x.CurrentDistance);
      
      foreach (var singleReindeer in reindeer.Where(x => x.CurrentDistance == topDistance))
        singleReindeer.Points++;
    }
    
    return reindeer.Max(x => x.Points);
  }

  private class ReindeerStats(int runningTime, int restTime, int speed)
  {
    public int CurrentDistance { get; private set; }
    public int Points { get; set; }
    
    private int RunningTime { get; } = runningTime;
    private int RestTime { get; } = restTime;
    private int Speed { get; } = speed;

    private int TotalRunTime => RunningTime + RestTime;

    public void ProcessDistanceForSecond(int second)
    {
      var secondInRun = ((second - 1) % TotalRunTime) + 1;
      if (secondInRun <= RunningTime)
        CurrentDistance += Speed;
    }
  }

  [GeneratedRegex(@"(?<name>\w+) can fly (?<speed>\d+) km/s for (?<runningTime>\d+) seconds, but then must rest for (?<restTime>\d+) seconds.")]
  private static partial Regex ReindeerStatsRegex();
}