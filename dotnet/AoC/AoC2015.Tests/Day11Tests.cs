namespace AoC2015.Tests;

public class Day11Tests
{
  [Test]
  [Arguments("abcdefgh", "abcdffaa")]
  [Arguments("ghijklmn", "ghjaabcc")]
  public async Task PartOneTest(string input, string expected)
  {
    // Arrange
    // Act
    var result = Day11.WithoutChannels([input]);
    
    // Assert
    await Assert.That(result).IsEqualTo(expected);
  }
}