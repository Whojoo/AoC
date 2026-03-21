using BenchmarkDotNet.Attributes;

using Shared;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day16Benchmark
{
  private readonly string[] _input = InputReader.ReadChallengeInput(2015, 16);
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day16.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return Day16.PartTwo(_input);
  }
}