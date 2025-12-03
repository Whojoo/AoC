using BenchmarkDotNet.Attributes;

using Shared;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day15Benchmark
{
  private readonly string[] _input = InputReader.ReadChallengeInput(2015, 15);
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day15.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return Day15.PartTwo(_input);
  }
}