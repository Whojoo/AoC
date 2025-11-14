namespace AoC2015.Tests;

public class Day14Tests
{
  [Fact]
  public void PartOneTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day14.txt");
    const int expectedResult = 1120;
    
    // Act
    var result = Day14.ReindeerRace(input, 1000);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
  
  [Fact]
  public void PartTwoTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day14.txt");
    const int expectedResult = 689;
    
    // Act
    var result = Day14.AdvancedReindeerRace(input, 1000);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
}