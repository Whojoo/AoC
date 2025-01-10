// See https://aka.ms/new-console-template for more information

using BenchmarkDotNet.Running;
using Dotnet;


// RunPuzzles();
RunBenchmark();

return;

void RunPuzzles()
{
    Day7.Run();
    Day5.Run();
}

void RunBenchmark()
{
    BenchmarkRunner.Run<AoCBenchmark>();
}