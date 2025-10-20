using BenchmarkDotNet.Attributes;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day10Benchmark
{
  private readonly string[] _input = File.ReadAllLines("../../../../../../../input/day10.txt");
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return Day10.PartOne(_input);
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return Day10.PartTwo(_input);
  }
}