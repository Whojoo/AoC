namespace Shared;

public static class InputReader
{
  private static readonly string InputPath = Environment.GetFolderPath(Environment.SpecialFolder.UserProfile);
  
  public static string[] ReadChallengeInput(int year, int day) => 
    ReadInputAsync(year, $"day{day}", "input").GetAwaiter().GetResult();
  
  public static Task<string[]> ReadTestInputAsync(int year, int day) =>
    ReadInputAsync(year, $"day{day}", "test-input");
  
  public static Task<string[]> ReadTestInputAsync(int year, int day, string suffix) =>
    ReadInputAsync(year, $"day{day}-{suffix}", "test-input");

  private static async Task<string[]> ReadInputAsync(int year, string fileName, string inputFolderName) =>
    await File.ReadAllLinesAsync($"{InputPath}/AdventOfCode/{inputFolderName}/{year}/{fileName}.txt");
}