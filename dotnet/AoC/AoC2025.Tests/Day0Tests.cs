using Shared;

namespace AoC2025.Tests;

public class Day0Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2025, 0);
    const int expectedResult = 0;
    
    // Act
    var result = Day0.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartTwoTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2025, 0);
    const int expectedResult = 0;
    
    // Act
    var result = Day0.PartTwo(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}