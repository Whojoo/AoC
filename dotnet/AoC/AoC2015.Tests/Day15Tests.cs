namespace AoC2015.Tests;

public class Day15Tests
{
  [Fact]
  public void PartOneTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day15.txt");
    const int expectedResult = 62842880;
    
    // Act
    var result = Day15.PartOne(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
  
  [Fact]
  public void PartTwoTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day15.txt");
    const int expectedResult = 57600000;
    
    // Act
    var result = Day15.PartTwo(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
}