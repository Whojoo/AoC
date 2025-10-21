using System.Runtime.ExceptionServices;
using System.Text.RegularExpressions;
using System.Threading.Channels;

namespace AoC2015;

public static partial class Day11
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day11.txt");

    var firstPassword = PartOne(lines).GetAwaiter().GetResult();
    Console.WriteLine($"Part one: {firstPassword}");
    Console.WriteLine($"Part two: {PartOne([firstPassword]).GetAwaiter().GetResult()}");
  }

  public static async Task<string> PartOne(string[] input)
  {
    var ctSource = new CancellationTokenSource();
    const bool log = false;
    var passwordChannel = Channel.CreateBounded<string>(new BoundedChannelOptions(100) { FullMode = BoundedChannelFullMode.Wait });
    var resultChannel = Channel.CreateUnbounded<string>();

    _ = Task.Run(async () => await GeneratePasswords(passwordChannel.Writer, input[0], log, ctSource.Token), ctSource.Token);
    _ = Task.Run(async () => await ValidatePasswords(passwordChannel.Reader, resultChannel.Writer, log, ctSource.Token), ctSource.Token);
    _ = Task.Run(async () => await ValidatePasswords(passwordChannel.Reader, resultChannel.Writer, log, ctSource.Token), ctSource.Token);

    var result = await resultChannel.Reader.ReadAsync(CancellationToken.None);
    await ctSource.CancelAsync();
    
    return result;
  }
  
  public static int PartTwo(string[] input)
  {
    return 0;
  }

  private static async Task GeneratePasswords(ChannelWriter<string> outputChannel, string startingPassword, bool log, CancellationToken ct)
  {
    var password = new string(startingPassword);
    while (!ct.IsCancellationRequested)
    {
      var chars = password.ToList();
      for (var i = chars.Count - 1; i >= 0; i--)
      {
        if (chars[i] == 'z')
        {
          chars[i] = 'a';
        }
        else
        {
          chars[i]++;
          break;
        }
      }

      password = string.Join(string.Empty, chars);
      if (log)
        Console.WriteLine($"Sending password {password}");
      await outputChannel.WriteAsync(password, ct);
    }
  }

  private static async Task ValidatePasswords(
    ChannelReader<string> inputChannel,
    ChannelWriter<string> outputChannel, 
    bool log,
    CancellationToken ct)
  {
    await foreach (var password in inputChannel.ReadAllAsync(ct))
    {
      if (log)
        Console.WriteLine($"Processing password {password}");

      if (ForbiddenCharacters().Match(password).Success)
      {
        if (log)
          Console.WriteLine($"Password {password} has a forbidden character");
        continue;
      }

      if (!HasTwoDoubles(password))
      {
        if (log)
          Console.WriteLine($"Password {password} does not have at least 2 doubles");
        continue;
      }

      if (!HasTwoIncrements(password))
      {
        if (log)
          Console.WriteLine($"Password {password} does not have at least 2 increments");
        continue;
      }
      
      await outputChannel.WriteAsync(password, CancellationToken.None);
    }
  }

  private static bool HasTwoDoubles(string password)
  {
    var foundFirstDouble = false;

    for (var i = 0; i < password.Length - 1; i++)
    {
      var character = password[i];
      var nextCharacter = password[i + 1];

      if (character != nextCharacter)
      {
        continue;
      }

      if (foundFirstDouble)
        return true;

      // No overlap, so give i an extra increment
      i++;
      foundFirstDouble = true;
    }

    return false;
  }

  private static bool HasTwoIncrements(string password)
  {
    for (var i = 0; i < password.Length - 2; i++)
    {
      var character = password[i];
      var secondCharacter = password[i + 1];
      var thirdCharacter = password[i + 2];

      if ((character + 1) == secondCharacter && (character + 2) == thirdCharacter)
      {
        return true;
      }
    }

    return false;
  }

  [GeneratedRegex("[iol]+")]
  private static partial Regex ForbiddenCharacters();
}