using BenchmarkDotNet.Attributes;

using Shared;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day8Benchmark
{
  private readonly string[] _input = InputReader.ReadChallengeInput(2015, 8);
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day8.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart1V2()
  {
    return Day8.PartOneV2(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return Day8.PartTwo(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2V2()
  {
    return Day8.PartTwoV2(_input);
  }
}