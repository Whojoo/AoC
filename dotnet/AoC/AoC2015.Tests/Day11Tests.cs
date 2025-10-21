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
    var result = await Day11.WithChannels([input]);
    
    // Assert
    Assert.Equal(expected, result);
  }
}