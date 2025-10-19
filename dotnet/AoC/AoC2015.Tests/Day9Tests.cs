namespace AoC2015.Tests;

public class Day9Tests
{
  [Fact]
  public void PartOneTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day9.txt");
    const int expectedResult = 605;
    
    // Act
    var result = Day9.PartOne(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
  
  [Fact]
  public void PartTwoTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day9.txt");
    const int expectedResult = 982;
    
    // Act
    var result = Day9.PartTwo(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
}