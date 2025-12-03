using BenchmarkDotNet.Attributes;

using Shared;

namespace AoC2025.Benchmarks;

[MemoryDiagnoser]
public class Day0Benchmark
{
  private readonly string[] _input = InputReader.ReadChallengeInput(2025, 0);
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day0.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return Day0.PartTwo(_input);
  }
}