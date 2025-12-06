using Shared;

namespace AoC2025.Tests;

public class Day2Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2025, 2);
    const ulong expectedResult = 1227775554;
    
    // Act
    var result = Day2.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }

  [Test]
  [Arguments("11-22", 33)]
  [Arguments("95-115", 99)]
  [Arguments("998-1012", 1010)]
  [Arguments("1188511880-1188511890", 1188511885)]
  [Arguments("222220-222224", 222222)]
  [Arguments("1698522-1698528", 0)]
  [Arguments("446443-446449", 446446)]
  [Arguments("38593856-38593862", 38593859)]
  public async Task PartOneMiniTests(string input, ulong expectedResult)
  {
    // Arrange & Act
    var result = Day2.PartOne([input]);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartTwoTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2025, 2);
    const ulong expectedResult = 4174379265;
    
    // Act
    var result = Day2.PartTwo(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  [Arguments("11-22", 11 + 22)]
  [Arguments("95-115", 99 + 111)]
  [Arguments("998-1012", 999 + 1010)]
  [Arguments("1188511880-1188511890", 1188511885)]
  [Arguments("222220-222224", 222222)]
  [Arguments("1698522-1698528", 0)]
  [Arguments("446443-446449", 446446)]
  [Arguments("38593856-38593862", 38593859)]
  [Arguments("565653-565659", 565656)]
  [Arguments("824824821-824824827", 824824824)]
  [Arguments("2121212118-2121212124", 2121212121)]
  public async Task PartTwoMiniTests(string input, ulong expectedResult)
  {
    // Arrange & Act
    var result = Day2.PartTwo([input]);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}