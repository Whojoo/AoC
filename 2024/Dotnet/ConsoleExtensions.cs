namespace Dotnet;
public static class ConsoleExtensions
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
