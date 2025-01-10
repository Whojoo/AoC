using BenchmarkDotNet.Attributes;

namespace Dotnet;

[MemoryDiagnoser]
[LongRunJob]
public class AoCBenchmark
{
    private List<string> _input = [];

    [IterationSetup(Targets = [nameof(Day7Part1), nameof(Day7Part2)])]
    public void SetupDay7()
    {
        _input = InputReader.ReadInput("../../../../../input/day7.txt");
    }

    [Benchmark]
    public long Day7Part1() => Day7.CalculatePart1(_input);

    [Benchmark]
    public long Day7Part2() => Day7.CalculatePart2(_input);

    [IterationSetup(Targets = [nameof(Day5Part1), nameof(Day5Part2)])]
    public void SetupDay5()
    {
        _input = InputReader.ReadInput("../../../../../input/day5.txt");
    }

    [Benchmark]
    public int Day5Part1() => Day5.CalculatePart1(_input);

    [Benchmark]
    public int Day5Part2() => Day5.CalculatePart2(_input);
}