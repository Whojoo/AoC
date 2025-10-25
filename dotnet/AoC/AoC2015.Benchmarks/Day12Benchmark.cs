using BenchmarkDotNet.Attributes;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day12Benchmark
{
  private readonly string[] _input = File.ReadAllLines("../../../../../../../input/day12.txt");
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day12.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart1Functional()
  {
    return Day12.PartOneFunctional(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return Day12.PartTwo(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2Functional()
  {
    return Day12.PartTwoFunctional(_input);
  }
}