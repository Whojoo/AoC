using BenchmarkDotNet.Attributes;

using Shared;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day10Benchmark
{
  private readonly string[] _input = InputReader.ReadChallengeInput(2015, 10);
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day10.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart1MoreBuilder()
  {
    return Day10.PartOneMoreBuilder(_input);
  }
}