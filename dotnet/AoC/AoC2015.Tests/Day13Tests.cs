namespace AoC2015.Tests;

public class Day13Tests
{
  [Fact]
  public void PartOneTest()
  {
    // Arrange
    var input = File.ReadAllLines("test-input/day13.txt");
    const int expectedResult = 330;
    
    // Act
    var result = Day13.PartOne(input);
    
    // Assert
    Assert.Equal(expectedResult, result);
  }
}