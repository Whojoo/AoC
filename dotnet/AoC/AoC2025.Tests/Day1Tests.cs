using Shared;

namespace AoC2025.Tests;

public class Day1Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2025, 1);
    const int expectedResult = 3;
    
    // Act
    var result = Day1.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartOneTestExtra()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2025, 1, "extra");
    const int expectedResult = 2;
    
    // Act
    var result = Day1.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartTwoTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2025, 1);
    const int expectedResult = 6;
    
    // Act
    var result = Day1.PartTwo(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}