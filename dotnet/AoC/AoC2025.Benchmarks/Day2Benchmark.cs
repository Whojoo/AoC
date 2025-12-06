using BenchmarkDotNet.Attributes;

using Shared;

namespace AoC2025.Benchmarks;

[MemoryDiagnoser]
public class Day2Benchmark
{
  private readonly string[] _input = InputReader.ReadChallengeInput(2025, 2);
  
  [Benchmark]
  public ulong BenchmarkPart1()
  {
    return Day2.PartOne(_input);
  }
  
  [Benchmark]
  public ulong BenchmarkPart2()
  {
    return Day2.PartTwo(_input);
  }
}