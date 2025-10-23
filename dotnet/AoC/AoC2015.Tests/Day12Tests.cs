namespace AoC2015.Tests;

public class Day12Tests
{
  [Theory]
  [InlineData("[1,2,3]", 6)]
  [InlineData("{\"a\":2,\"b\":4}", 6)]
  [InlineData("[[[3]]]", 3)]
  [InlineData("{\"a\":{\"b\":4},\"c\":-1}", 3)]
  [InlineData("{\"a\":[-1,1]}", 0)]
  [InlineData("[-1,{\"a\":1}] ", 0)]
  [InlineData("[]", 0)]
  [InlineData("{}", 0)]
  public void PartOneTest(string input, int expected)
  {
    // Arrange
    // Act
    var result = Day12.PartOne([input]);
    
    // Assert
    Assert.Equal(expected, result);
  }
  
  [Theory]
  [InlineData("[1,2,3]", 6)]
  [InlineData("{\"a\":2,\"b\":4}", 6)]
  [InlineData("[[[3]]]", 3)]
  [InlineData("{\"a\":{\"b\":4},\"c\":-1}", 3)]
  [InlineData("{\"a\":[-1,1]}", 0)]
  [InlineData("[-1,{\"a\":1}] ", 0)]
  [InlineData("[]", 0)]
  [InlineData("{}", 0)]
  [InlineData("[1,{\"c\":\"red\",\"b\":2},3]", 4)]
  [InlineData("{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", 0)]
  [InlineData("[1,\"red\",5]", 6)]
  public void PartTwoTest(string input, int expected)
  {
    // Arrange
    // Act
    var result = Day12.PartTwo([input]);
    
    // Assert
    Assert.Equal(expected, result);
  }
}