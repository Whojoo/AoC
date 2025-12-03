using Shared;

namespace AoC2015.Tests;

public class Day9Tests
{
  [Test]
  public async Task PartOneTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 9);
    const int expectedResult = 605;
    
    // Act
    var result = Day9.PartOne(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
  
  [Test]
  public async Task PartTwoTest()
  {
    // Arrange
    var input = await InputReader.ReadTestInputAsync(2015, 9);
    const int expectedResult = 982;
    
    // Act
    var result = Day9.PartTwo(input);
    
    // Assert
    await Assert.That(result).IsEqualTo(expectedResult);
  }
}