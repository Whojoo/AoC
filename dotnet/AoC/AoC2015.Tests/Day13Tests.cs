using Shared;

namespace AoC2015.Tests;

public class Day13Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 13);
    const int expectedResult = 330;
    
    // Act
    var result = Day13.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}