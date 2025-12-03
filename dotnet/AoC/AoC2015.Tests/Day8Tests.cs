using Shared;

namespace AoC2015.Tests;

public class Day8Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 8);
    const int expectedResult = 15;
    
    // Act
    var result = Day8.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }

  [Test]
  [Arguments("\"\"", 2)]                         // ""
  [Arguments("\"\\\"\"", 3)]                     // "\""
  [Arguments("\"\\\\\"", 3)]                     // "\\"
  [Arguments("\"\\x24\"", 5)]                    // "\x24"
  [Arguments("\"ecn\\x50ooprbstnq\"", 5)]        // "ecn\x50ooprbstnq"
  // Ignore the hex on \\x
  [Arguments("\"x\\\"\\xcaj\\\\xwwvpdldz\"", 7)] // "x\"\xcaj\\xwwvpdldz"
  // Do not ignore hex on \\\x, but also do not double count the \\
  [Arguments("\"fdan\\\\\\x9e\"", 6)]            // "fdan\\\x9e"
  public async Task PartOneTestEscapes(string input, int expectedResult)
  {
    // Arrange
    // Act
    var result = Day8.PartOne([input]);

    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartTwoTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 8);
    const int expectedResult = 25;
    
    // Act
    var result = Day8.PartTwo(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }

  [Test]
  [Arguments("\"\"", 4)]                         // "" -> "\"\""
  [Arguments("\"\\\"\"", 6)]                     // "\"" -> "\"\\\"\""
  [Arguments("\"\\\\\"", 6)]                     // "\\" -> "\"\\\\\""
  [Arguments("\"\\x24\"", 5)]                    // "\x24" -> "\"\\x24\""
  // [Arguments("\"ecn\\x50ooprbstnq\"", 5)]        // "ecn\x50ooprbstnq"
  // [Arguments("\"x\\\"\\xcaj\\\\xwwvpdldz\"", 7)] // "x\"\xcaj\\xwwvpdldz"
  // [Arguments("\"fdan\\\\\\x9e\"", 6)]            // "fdan\\\x9e"
  public async Task PartTwoTestSeparateTestInput(string input, int expectedResult)
  {
    // Arrange
    // Act
    var result = Day8.PartTwo([input]);

    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}