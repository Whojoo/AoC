using System.Diagnostics;
using System.Text.RegularExpressions;

namespace Dotnet;

public static partial class Day14
{
    public static void Run()
    {
        var totalStartTime = Stopwatch.GetTimestamp();
        var input = InputReader.ReadInput("day14.txt");

        var startTime = Stopwatch.GetTimestamp();
        var part1 = Part1(input, 101, 103);
        var part1Time = Stopwatch.GetElapsedTime(startTime);

        // startTime = Stopwatch.GetTimestamp();
        // var part2 = Part2(input);
        // var part2Time = Stopwatch.GetElapsedTime(startTime);

        var totalElapsed = Stopwatch.GetElapsedTime(totalStartTime);

        Console.WriteLine("Day 14");
        Console.WriteLine($"Part 1: {part1} in {part1Time.TotalMilliseconds} ms");
        // Console.WriteLine($"Part 2: {part2} in {part2Time.TotalMilliseconds} ms");
        Console.WriteLine($"Total elapsed (including reads): {totalElapsed.TotalMilliseconds} ms");
        Console.WriteLine();
    }

    public static long Part1(List<string> input, int width, int height)
    {
        const int secondsToProcess = 100;
        
        var robots = ParseInputToRobots(input);

        foreach (var robot in robots)
        {
            var totalVelocity = robot.Velocity.GetScaledVector(secondsToProcess);
            robot.Position.Add(totalVelocity);
            robot.Position.LimitByBorders(width, height, robot.Velocity.X < 0, robot.Velocity.Y < 0);
        }
        
        var ignoredX = (width - 1) / 2;
        var ignoredY = (height - 1) / 2;

        var relevantRobots = robots.Where(robot => robot.Position.X != ignoredX && robot.Position.Y != ignoredY);
        var topLeft = 0;
        var bottomLeft = 0;
        var topRight = 0;
        var bottomRight = 0;

        foreach (var robot in relevantRobots)
        {
            if (robot.Position.X < ignoredX)
            {
                if (robot.Position.Y < ignoredY)
                    topLeft++;
                else
                    bottomLeft++;
            }
            else
            {
                if (robot.Position.Y < ignoredY)
                    topRight++;
                else
                    bottomRight++;
            }
        }

        return topLeft * bottomRight * topRight * bottomLeft;
    }

    public static long Part2(List<string> input)
    {
        var robots = ParseInputToRobots(input);
        const int width = 101;
        const int height = 103;

        const int start = 200_000;
        const int batch = 50_000;
        var seconds = start;
        File.Delete("day14-pt2-shizzle.txt");

        foreach (var robot in robots)
        {
            robot.Position.Add(robot.Velocity.GetScaledVector(seconds));
            robot.Position.LimitByBorders(width, height, robot.Velocity.X < 0, robot.Velocity.Y < 0);
        }
        
        while (seconds < start + batch)
        {
            seconds++;
            
            foreach (var robot in robots)
            {
                robot.Position.Add(robot.Velocity);
                robot.Position.LimitByBorders(width, height, robot.Velocity.X < 0, robot.Velocity.Y < 0);
            }
            
            var world = new char[width][];

            for(var rowIndex = 0; rowIndex < world.Length; rowIndex++)
            {
                world[rowIndex] = new char[height];

                var relevantRobots = robots.Where(robot => robot.Position.X == rowIndex).ToArray();
                
                for (var colIndex = 0; colIndex < world[rowIndex].Length; colIndex++)
                {
                    var robot = relevantRobots.FirstOrDefault(robot => robot.Position.Y == colIndex);
                    world[rowIndex][colIndex] = robot is null ? '.' : '#';
                }
            }

            var lines = new List<string>(world.Length + 1);
            lines.Add($"Seconds: {seconds}");
            foreach (var row in world)
            {
                lines.Add(new string(row));
            }
            
            File.AppendAllLines("day14-pt2-shizzle.txt", lines);
        }
        
        return seconds;
    }

    [GeneratedRegex(@"p=(?<PosX>\d+),(?<PosY>\d+) v=(?<VelX>-?\d+),(?<VelY>-?\d+)")]
    private static partial Regex RobotParameterRegex();

    private static Robot[] ParseInputToRobots(List<string> input) =>
        input
            .Select(x =>
            {
                var match = RobotParameterRegex().Match(x);
                var posX = int.Parse(match.Groups["PosX"].Value);
                var posY = int.Parse(match.Groups["PosY"].Value);
                var velX = int.Parse(match.Groups["VelX"].Value);
                var velY = int.Parse(match.Groups["VelY"].Value);

                return new Robot(new Vector(posX, posY), new Vector(velX, velY));
            })
            .ToArray();

    private record Robot(Vector Position, Vector Velocity);

    private class Vector(int x, int y) : IEquatable<Vector>
    {
        public int X { get; private set; } = x;
        public int Y { get; private set; } = y;

        public void Add(Vector vector)
        {
            X += vector.X;
            Y += vector.Y;
        }

        public void LimitByBorders(int width, int height, bool reversedWidth, bool reversedHeight)
        {
            X = Math.Abs(X) % width;
            Y = Math.Abs(Y) % height;
            
            if (reversedWidth && X != 0) X = width - X;
            if (reversedHeight && Y != 0) Y = height - Y;
        }

        public Vector GetScaledVector(int scalar) => new(X * scalar, Y * scalar);

        public override string ToString() => $"({X}, {Y})";

        public bool Equals(Vector? other)
        {
            if (other is null)
            {
                return false;
            }

            if (ReferenceEquals(this, other))
            {
                return true;
            }

            return X == other.X && Y == other.Y;
        }

        public override bool Equals(object? obj)
        {
            if (obj is null)
            {
                return false;
            }

            if (ReferenceEquals(this, obj))
            {
                return true;
            }

            if (obj.GetType() != GetType())
            {
                return false;
            }

            return Equals((Vector)obj);
        }

        public override int GetHashCode()
        {
            return HashCode.Combine(X, Y);
        }

        public static bool operator ==(Vector? left, Vector? right)
        {
            return Equals(left, right);
        }

        public static bool operator !=(Vector? left, Vector? right)
        {
            return !Equals(left, right);
        }
    }
}