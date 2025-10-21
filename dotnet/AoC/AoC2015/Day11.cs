using System.Text.RegularExpressions;
using System.Threading.Channels;

namespace AoC2015;

public static partial class Day11
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day11.txt");

    var firstPassword = WithChannels(lines).GetAwaiter().GetResult();
    Console.WriteLine($"Part one: {firstPassword}");
    Console.WriteLine($"Part two: {WithChannels([firstPassword]).GetAwaiter().GetResult()}");

    var password = WithoutChannels(lines);
    Console.WriteLine($"Part one: {password}");
    Console.WriteLine($"Part two: {WithoutChannels([password])}");
  }

  public static async Task<string> WithChannels(string[] input)
  {
    var ctSource = new CancellationTokenSource();
    var passwordChannel = Channel.CreateBounded<string>(new BoundedChannelOptions(100) { FullMode = BoundedChannelFullMode.Wait });
    var resultChannel = Channel.CreateUnbounded<string>();

    _ = Task.Run(async () => await GeneratePasswords(passwordChannel.Writer, input[0], ctSource.Token), ctSource.Token);
    _ = Task.Run(async () => await ValidatePasswords(passwordChannel.Reader, resultChannel.Writer, ctSource.Token), ctSource.Token);
    _ = Task.Run(async () => await ValidatePasswords(passwordChannel.Reader, resultChannel.Writer, ctSource.Token), ctSource.Token);

    var result = await resultChannel.Reader.ReadAsync(CancellationToken.None);
    await ctSource.CancelAsync();
    
    return result;
  }
  
  public static string WithoutChannels(string[] input)
  {
    var password = input[0];

    while (true)
    {
      password = GenerateNewPassword(password);
      if (ValidatePassword(password))
        break;
    }

    return password;
  }

  private static async Task GeneratePasswords(ChannelWriter<string> outputChannel, string startingPassword, CancellationToken ct)
  {
    var password = new string(startingPassword);
    while (!ct.IsCancellationRequested)
    {
      password = GenerateNewPassword(password);
      await outputChannel.WriteAsync(password, ct);
    }
  }

  private static string GenerateNewPassword(string password)
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

    return string.Join(string.Empty, chars);
  }

  private static async Task ValidatePasswords(
    ChannelReader<string> inputChannel,
    ChannelWriter<string> outputChannel, 
    CancellationToken ct)
  {
    await foreach (var password in inputChannel.ReadAllAsync(ct))
    {
      if (!ValidatePassword(password))
        continue;
      
      await outputChannel.WriteAsync(password, CancellationToken.None);
    }
  }

  private static bool ValidatePassword(string password)
  {
    if (ForbiddenCharacters().Match(password).Success)
    {
      return false;
    }

    if (!HasTwoDoubles(password))
    {
      return false;
    }

    if (!HasTwoIncrements(password))
    {
      return false;
    }

    return true;
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