using BenchmarkDotNet.Attributes;

using Shared;

namespace AoC2025.Benchmarks;

[MemoryDiagnoser]
public class Day1Benchmark
{
  private readonly string[] _input = InputReader.ReadChallengeInput(2025, 1);
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day1.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return Day1.PartTwo(_input);
  }
}