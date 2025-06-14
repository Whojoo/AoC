namespace Dotnet.Test;

public static class TestInputReader
{
    public static List<string> ReadInput(string fileName)
    {
        return File.ReadAllLines($"testInput/{fileName}").ToList();
    }
}