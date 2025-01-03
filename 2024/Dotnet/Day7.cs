using System.Collections.Concurrent;
using System.Diagnostics;

namespace Dotnet;

public static class Day7
{
    private static readonly Dictionary<int, List<List<IOperation>>> OperationOptionsCache = [];
    private static readonly ConcurrentDictionary<int, List<List<IOperation>>> OperationOptionsCache2 = [];

    public static void Run()
    {
        var totalStartTime = Stopwatch.GetTimestamp();
        var input = InputReader
            .ReadInput("day7.txt")
            .Select(MapToEquation)
            .ToList();

        var startTime = Stopwatch.GetTimestamp();
        var part1 = CalculatePart1(input);
        var part1Time = Stopwatch.GetElapsedTime(startTime);

        startTime = Stopwatch.GetTimestamp();
        var part2 = CalculatePart2(input);
        var part2Time = Stopwatch.GetElapsedTime(startTime);

        var totalElapsed = Stopwatch.GetElapsedTime(totalStartTime);

        Console.WriteLine("Day 7");
        Console.WriteLine($"Part 1: {part1} in {part1Time.TotalMilliseconds} ms");
        Console.WriteLine($"Part 2: {part2} in {part2Time.TotalMilliseconds} ms");
        Console.WriteLine($"Total elapsed (including reads): {totalElapsed.TotalMilliseconds} ms");
        Console.WriteLine();
    }

    public static long CalculatePart1(List<Equation> input) =>
        input
            .Where(IsValidEquation)
            .Sum(equation => equation.TestValue);

    public static long CalculatePart2(List<Equation> input) => 
        input
            .AsParallel()
            .Where(IsValidEquation2)
            .Sum(equation => equation.TestValue);

    private static bool IsValidEquation(Equation equation) =>
        GetOperationOptions(equation.Numbers.Count)
            .Any(operations => IsValidEquationForOperations(equation, operations));
    
    private static bool IsValidEquation2(Equation equation) =>
        GetOperationOptions2(equation.Numbers.Count)
            .AsParallel()
            .Any(operations => IsValidEquationForOperations(equation, operations));

    private static bool IsValidEquationForOperations(Equation equation, List<IOperation> operations)
    {
        var number = equation.Numbers[0];
        for (var i = 1; i < equation.Numbers.Count; i++)
        {
            var operation = operations[i - 1];
            number = operation.Solve(number, equation.Numbers[i]);
        }

        return number == equation.TestValue;
    }

    public static Equation MapToEquation(string input)
    {
        var splitInput = input.Split(": ");
        var (testString, numbers) = (splitInput[0], splitInput[1]);

        return new Equation(
            long.Parse(testString),
            numbers.Split(" ").Select(long.Parse).ToList());
    }

    private static List<List<IOperation>> GetOperationOptions(int amountOfNumbers)
    {
        if (OperationOptionsCache.TryGetValue(amountOfNumbers, out var result))
        {
            return result;
        }

        var (add, multiply) = (new AddOperation(), new MultiplyOperation());
        List<List<IOperation>> operations = [[add], [multiply]];

        for (var i = 3; i <= amountOfNumbers; i++)
        {
            var newOperations = new List<List<IOperation>>(operations.Count * 2);
            newOperations.AddRange(operations.Select(operation => (List<IOperation>) [.. operation, add]));
            newOperations.AddRange(operations.Select(operation => (List<IOperation>) [.. operation, multiply]));
            operations = newOperations;
        }

        OperationOptionsCache[amountOfNumbers] = operations;
        return operations;
    }

    private static List<List<IOperation>> GetOperationOptions2(int amountOfNumbers)
    {
        if (OperationOptionsCache2.TryGetValue(amountOfNumbers, out var result))
        {
            return result;
        }

        var (add, multiply, concatenate) = (new AddOperation(), new MultiplyOperation(), new ConcatenateOperation());
        List<List<IOperation>> operations = [[add], [multiply], [concatenate]];

        for (var i = 3; i <= amountOfNumbers; i++)
        {
            var newOperations = new List<List<IOperation>>();
            newOperations.AddRange(operations.Select(operation => (List<IOperation>) [.. operation, add]));
            newOperations.AddRange(operations.Select(operation => (List<IOperation>) [.. operation, multiply]));
            newOperations.AddRange(operations.Select(operation => (List<IOperation>) [.. operation, concatenate]));
            operations = newOperations;
        }

        OperationOptionsCache2[amountOfNumbers] = operations;
        return operations;
    }

    public class Equation(long testValue, List<long> numbers)
    {
        public long TestValue { get; } = testValue;
        public List<long> Numbers { get; } = numbers;
    }

    public interface IOperation
    {
        long Solve(long left, long right);
    }

    public class AddOperation : IOperation
    {
        public long Solve(long left, long right) => left + right;
    }

    public class MultiplyOperation : IOperation
    {
        public long Solve(long left, long right) => left * right;
    }

    public class ConcatenateOperation : IOperation
    {
        public long Solve(long left, long right)
        {
            var multiplier = 10;
            while (multiplier <= right) multiplier *= 10;
            
            return left * multiplier + right;
        }
    }
}