using Shared;

namespace AoC2015.Tests;

public class Day14Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 14);
    const int expectedResult = 1120;
    
    // Act
    var result = Day14.ReindeerRace(input, 1000);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartTwoTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 14);
    const int expectedResult = 689;
    
    // Act
    var result = Day14.AdvancedReindeerRace(input, 1000);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}