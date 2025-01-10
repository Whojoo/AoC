using System.Collections.Concurrent;
using System.Diagnostics;

namespace Dotnet;

public static class Day7
{
    public static void Run()
    {
        var totalStartTime = Stopwatch.GetTimestamp();
        var input = InputReader.ReadInput("day7.txt");

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

    public static long CalculatePart1(List<string> input)
    {
        List<IOperation> operations =
        [
            new AddOperation(),
            new MultiplyOperation()
        ];

        return input
            .Select(MapToEquation)
            .Where(equation => IsValidEquationForOperations(equation, operations))
            .Sum(equation => equation.TestValue);
    }

    public static long CalculatePart2(List<string> input)
    {
        List<IOperation> operations =
        [
            new AddOperation(),
            new MultiplyOperation(),
            new ConcatenateOperation()
        ];

        return input
            .Select(MapToEquation)
            .Where(equation => IsValidEquationForOperations(equation, operations))
            .Sum(equation => equation.TestValue);
    }
    
    private static bool IsValidEquationForOperations(Equation equation, List<IOperation> operations)
    {
        var operationNodeStack = new Stack<OperationNode>();
        operationNodeStack.Push(new OperationNode(equation.TestValue, equation.Numbers));
        var firstNumber = equation.Numbers[0];

        do
        {
            var node = operationNodeStack.Pop();
            var numbers = node.NumbersLeft.ToList();
            var nextNumber = numbers.Last();

            var nextTotals = operations
                .Where(operation => operation.IsApplicable(node.CurrentTotal, nextNumber))
                .Select(operation => operation.PerformOperation(node.CurrentTotal, nextNumber))
                .ToArray();

            switch (numbers.Count)
            {
                case 2 when nextTotals.Any(number => number == firstNumber):
                    return true;
                case 2:
                    continue;
            }

            var nextNumbers = numbers.Take(numbers.Count - 1);
            foreach (var total in nextTotals)
            {
                // Multiple is accepted for the chance of not having to create a new list
                // ReSharper disable once PossibleMultipleEnumeration
                operationNodeStack.Push(new OperationNode(total, nextNumbers));
            }

        } while (operationNodeStack.Count > 0);

        return false;
    }

    private static Equation MapToEquation(string input)
    {
        var splitInput = input.Split(": ");
        var (testString, numbers) = (splitInput[0], splitInput[1]);

        return new Equation(
            long.Parse(testString),
            numbers.Split(" ").Select(long.Parse).ToList());
    }

    public class Equation(long testValue, List<long> numbers)
    {
        public long TestValue { get; } = testValue;
        public List<long> Numbers { get; } = numbers;
    }

    public interface IOperation
    {
        bool IsApplicable(long currentTotal, long nextNumber);
        long PerformOperation(long currentTotal, long nextNumber);
    }

    public class AddOperation : IOperation
    {
        public bool IsApplicable(long currentTotal, long nextNumber) => currentTotal >= nextNumber;
        public long PerformOperation(long currentTotal, long nextNumber) => currentTotal - nextNumber;
    }

    public class MultiplyOperation : IOperation
    {
        public bool IsApplicable(long currentTotal, long nextNumber) => currentTotal % nextNumber == 0;
        public long PerformOperation(long currentTotal, long nextNumber) => currentTotal / nextNumber;
    }

    public class ConcatenateOperation : IOperation
    {
        public bool IsApplicable(long currentTotal, long nextNumber)
        {
            var comparisonNumberFor = GetComparisonNumberFor(nextNumber);
            var totalWithoutNextNumber = currentTotal - nextNumber;
            return totalWithoutNextNumber > 0 && totalWithoutNextNumber % comparisonNumberFor == 0;
        }

        public long PerformOperation(long currentTotal, long nextNumber)
        {
            var comparisonNumberFor = GetComparisonNumberFor(nextNumber);
            var totalWithoutNextNumber = currentTotal - nextNumber;
            return totalWithoutNextNumber / comparisonNumberFor;
        }

        private static int GetComparisonNumberFor(long nextNumber)
        {
            var comparison = 10;
            while (comparison <= nextNumber) comparison *= 10;
            return comparison;
        }
    }

    public class OperationNode(long currentTotal, IEnumerable<long> numbersLeft)
    {
        public long CurrentTotal { get; } = currentTotal;
        public IEnumerable<long> NumbersLeft { get; } = numbersLeft;
    }
}