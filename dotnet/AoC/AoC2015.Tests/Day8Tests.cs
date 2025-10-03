namespace AoC2015.Tests;

public class Day8Tests
{
  [Fact]
  public void PartOneTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day8.txt");
    const int expectedResult = 15;
    
    // Act
    var result = Day8.PartOne(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }

  [Theory]
  [InlineData("\"\"", 2)]                         // ""
  [InlineData("\"\\\"\"", 3)]                     // "\""
  [InlineData("\"\\\\\"", 3)]                     // "\\"
  [InlineData("\"\\x24\"", 5)]                    // "\x24"
  [InlineData("\"ecn\\x50ooprbstnq\"", 5)]        // "ecn\x50ooprbstnq"
  // Ignore the hex on \\x
  [InlineData("\"x\\\"\\xcaj\\\\xwwvpdldz\"", 7)] // "x\"\xcaj\\xwwvpdldz"
  // Do not ignore hex on \\\x, but also do not double count the \\
  [InlineData("\"fdan\\\\\\x9e\"", 6)]            // "fdan\\\x9e"
  public void PartOneTestEscapes(string input, int expectedResult)
  {
    // Arrange
    // Act
    var result = Day8.PartOne([input]);

    // Assert
    Assert.Equal(expectedResult, result);
  }
  
  [Fact]
  public void PartTwoTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day8.txt");
    const int expectedResult = 25;
    
    // Act
    var result = Day8.PartTwo(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }

  [Theory]
  [InlineData("\"\"", 4)]                         // "" -> "\"\""
  [InlineData("\"\\\"\"", 6)]                     // "\"" -> "\"\\\"\""
  [InlineData("\"\\\\\"", 6)]                     // "\\" -> "\"\\\\\""
  [InlineData("\"\\x24\"", 5)]                    // "\x24" -> "\"\\x24\""
  // [InlineData("\"ecn\\x50ooprbstnq\"", 5)]        // "ecn\x50ooprbstnq"
  // [InlineData("\"x\\\"\\xcaj\\\\xwwvpdldz\"", 7)] // "x\"\xcaj\\xwwvpdldz"
  // [InlineData("\"fdan\\\\\\x9e\"", 6)]            // "fdan\\\x9e"
  public void PartTwoTestSeparateTestInput(string input, int expectedResult)
  {
    // Arrange
    // Act
    var result = Day8.PartTwo([input]);

    // Assert
    Assert.Equal(expectedResult, result);
  }
}