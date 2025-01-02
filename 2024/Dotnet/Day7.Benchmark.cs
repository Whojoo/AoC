using BenchmarkDotNet.Attributes;

namespace Dotnet;

[MemoryDiagnoser(false)]
public class Day7_Benchmark
{
    private List<Day7.Equation> _input = null!;

    [GlobalSetup]
    public void Setup()
    {
        _input = InputReader
            .ReadInput("../../../../../input/day7.txt")
            .Select(Day7.MapToEquation)
            .ToList();
    }
    
    [Benchmark]
    public long Part1() => Day7.CalculatePart1(_input);
    
    [Benchmark]
    public long Part2() => Day7.CalculatePart2(_input);
}