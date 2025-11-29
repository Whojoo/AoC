using BenchmarkDotNet.Attributes;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day15Benchmark
{
  private readonly string[] _input = File.ReadAllLines("../../../../../../../input/day15.txt");
  
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