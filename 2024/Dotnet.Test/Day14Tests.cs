using Shouldly;

namespace Dotnet.Test;

public class Day14Tests
{
    [Theory]
    [InlineData("day14-default.txt", 12)]
    public void Part1(string fileName, long expectedResult)
    {
        var input = TestInputReader.ReadInput(fileName);
        var result = Day14.Part1(input, 11, 7);
        result.ShouldBe(expectedResult);
    }
}