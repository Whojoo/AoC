using System.Security.Cryptography;
using System.Text;

using Microsoft.CSharp.RuntimeBinder;

namespace AoC2015;

public static class Day4
{
  public static void Run()
  {
    var line = File.ReadAllLines("input/day4.txt")[0];

    Console.WriteLine($"Part one: {PartOne(line)}");
    Console.WriteLine($"Part two: {PartTwo(line)}");
  }

  public static int PartOne(string hashSecret)
  {
    foreach (var index in Enumerable.Range(0, 1000000000))
    {
      var secret = $"{hashSecret}{index}";
      var hash = MD5.HashData(Encoding.UTF8.GetBytes(secret));
      var hex = Convert.ToHexString(hash);

      if (hex.StartsWith("00000"))
        return index;
    }

    return -1;
  }

  public static int PartTwo(string hashSecret)
  {
    foreach (var index in Enumerable.Range(0, 1000000000))
    {
      var secret = $"{hashSecret}{index}";
      var hash = MD5.HashData(Encoding.UTF8.GetBytes(secret));
      var hex = Convert.ToHexString(hash);

      if (hex.StartsWith("000000"))
        return index;
    }

    return -1;
  }
}