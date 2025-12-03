using Shared;

namespace AoC2015.Tests;

public class Day15Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 15);
    const int expectedResult = 62842880;
    
    // Act
    var result = Day15.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartTwoTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 15);
    const int expectedResult = 57600000;
    
    // Act
    var result = Day15.PartTwo(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}