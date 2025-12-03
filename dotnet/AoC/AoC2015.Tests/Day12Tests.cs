namespace AoC2015.Tests;

public class Day12Tests
{
  [Test]
  [Arguments("[1,2,3]", 6)]
  [Arguments("{\"a\":2,\"b\":4}", 6)]
  [Arguments("[[[3]]]", 3)]
  [Arguments("{\"a\":{\"b\":4},\"c\":-1}", 3)]
  [Arguments("{\"a\":[-1,1]}", 0)]
  [Arguments("[-1,{\"a\":1}] ", 0)]
  [Arguments("[]", 0)]
  [Arguments("{}", 0)]
  public async Task PartOneTest(string input, int expected)
  {
    // Arrange
    // Act
    var result = Day12.PartOne([input]);
    
    // Assert
    await Assert.That(result).IsEqualTo(expected);
  }
  
  [Test]
  [Arguments("[1,2,3]", 6)]
  [Arguments("{\"a\":2,\"b\":4}", 6)]
  [Arguments("[[[3]]]", 3)]
  [Arguments("{\"a\":{\"b\":4},\"c\":-1}", 3)]
  [Arguments("{\"a\":[-1,1]}", 0)]
  [Arguments("[-1,{\"a\":1}] ", 0)]
  [Arguments("[]", 0)]
  [Arguments("{}", 0)]
  [Arguments("[1,{\"c\":\"red\",\"b\":2},3]", 4)]
  [Arguments("{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", 0)]
  [Arguments("[1,\"red\",5]", 6)]
  public async Task PartTwoTest(string input, int expected)
  {
    // Arrange
    // Act
    var result = Day12.PartTwo([input]);
    
    // Assert
    await Assert.That(result).IsEqualTo(expected);
  }
}