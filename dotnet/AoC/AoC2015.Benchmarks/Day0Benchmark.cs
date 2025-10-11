using BenchmarkDotNet.Attributes;

namespace AoC2015.Benchmarks;

[MemoryDiagnoser]
public class Day0Benchmark
{
  private readonly string[] _input = File.ReadAllLines("../../../../../../../input/day0.txt");
  
  [Benchmark]
  public int BenchmarkPart1()
  {
    return 0;
  }
  
  [Benchmark]
  public int BenchmarkPart2()
  {
    return 0;
  }
}