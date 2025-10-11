using BenchmarkDotNet.Attributes;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day8Benchmark
{
  private readonly string[] _input = File.ReadAllLines("../../../../../../../input/day8.txt");
  
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