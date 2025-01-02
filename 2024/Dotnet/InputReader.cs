namespace Dotnet;

public static class InputReader
{
    public static List<string> ReadInput(string fileName)
    {
        return File.ReadAllLines($"../../../../input/{fileName}").ToList();
    }
}