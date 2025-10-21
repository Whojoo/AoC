namespace AoC2015.Tests;

public class Day11Tests
{
  [Theory]
  [InlineData("abcdefgh", "abcdffaa")]
  [InlineData("ghijklmn", "ghjaabcc")]
  public async Task PartOneTest(string input, string expected)
  {
    // Arrange
    // Act
    var result = await Day11.PartOne([input]);
    
    // Assert
    Assert.Equal(expected, result);
  }
  
  [Fact]
  public void PartTwoTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day11.txt");
    const int expectedResult = 25;
    
    // Act
    var result = Day11.PartTwo(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
}