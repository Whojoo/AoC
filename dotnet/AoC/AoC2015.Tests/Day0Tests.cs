namespace AoC2015.Tests;

public class Day0Tests
{
  [Fact]
  public void PartOneTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day0.txt");
    const int expectedResult = 15;
    
    // Act
    var result = Day0.PartOne(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
  
  [Fact]
  public void PartTwoTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day0.txt");
    const int expectedResult = 25;
    
    // Act
    var result = Day0.PartTwo(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
}