// See https://aka.ms/new-console-template for more information

using AoC2024.Day5;

var input = Console.In.ReadInput();

var first = Handler.CalculatePart1(input);
var second = Handler.CalculatePart2(input);

Console.WriteLine(first);
Console.WriteLine(second);

internal static class ConsoleExtensions
{
    public static List<string> ReadInput(this TextReader input)
    {
        List<string> output = [];

        var whitespaceAllowed = true;

        while (input.ReadLine() is { } line)
        {
            if (string.IsNullOrWhiteSpace(line) && !whitespaceAllowed)
            {
                break;
            }

            if (string.IsNullOrWhiteSpace(line))
            {
                whitespaceAllowed = false;
            }
            
            Console.WriteLine(line);
            output.Add(line);
        }

        return output;
    }
}