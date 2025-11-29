using System.Text.RegularExpressions;

using Shared;

namespace AoC2015;

public static partial class Day15
{
  public static void Run()
  {
    var lines = File.ReadAllLines("input/day15.txt");

    Console.WriteLine($"Part one: {PartOne(lines)}");
    Console.WriteLine($"Part two: {PartTwo(lines)}");
  }

  public static int PartOne(string[] input)
  {
    var ingredients = input
      .Select(MapLineToIngredient)
      .ToArray();

    return PartOneRecursion(ingredients, 0, 100)
      .Where(x => x is { Capacity: >= 0, Durability: >= 0, Flavor: >= 0, Texture: >= 0 })
      .Max(x => x.Capacity * x.Durability * x.Flavor * x.Texture);
  }

  public static int PartTwo(string[] input)
  {
    var ingredients = input
      .Select(MapLineToIngredient)
      .ToArray();

    return PartOneRecursion(ingredients, 0, 100)
      .Where(x => x is { Capacity: >= 0, Durability: >= 0, Flavor: >= 0, Texture: >= 0, Calories: 500 })
      .Max(x => x.Capacity * x.Durability * x.Flavor * x.Texture);
  }
  
  private static Ingredient MapLineToIngredient(string x)
  {
    var match = IngredientRegex().Match(x);
    return new Ingredient(match.Groups["capacity"].IntValue(),
      match.Groups["durability"].IntValue(),
      match.Groups["flavor"].IntValue(),
      match.Groups["texture"].IntValue(),
      match.Groups["calories"].IntValue());
  }

  private static IEnumerable<IngredientCalculation> PartOneRecursion(Ingredient[] ingredients, int currentIngredient,
    int ingredientsLeftOver)
  {
    if (currentIngredient >= ingredients.Length) yield break;

    var ingredient = ingredients[currentIngredient];
    
    if (currentIngredient == ingredients.Length - 1)
    {
      yield return new IngredientCalculation(
        ingredient.Capacity * ingredientsLeftOver,
        ingredient.Durability * ingredientsLeftOver,
        ingredient.Flavor * ingredientsLeftOver,
        ingredient.Texture * ingredientsLeftOver,
        ingredient.Calories * ingredientsLeftOver);
      yield break;
    }

    foreach (var ingredientsUsedThisIteration in Enumerable.Range(0, ingredientsLeftOver))
    {
      var indexedCalculation = new IngredientCalculation(
        ingredient.Capacity * ingredientsUsedThisIteration,
        ingredient.Durability * ingredientsUsedThisIteration,
        ingredient.Flavor * ingredientsUsedThisIteration,
        ingredient.Texture * ingredientsUsedThisIteration,
        ingredient.Calories * ingredientsUsedThisIteration);

      var ingredientsForNextIteration = ingredientsLeftOver - ingredientsUsedThisIteration;

      foreach (var ingredientCalculation in PartOneRecursion(ingredients, currentIngredient + 1, ingredientsForNextIteration))
      {
        yield return new IngredientCalculation(
          Capacity: ingredientCalculation.Capacity + indexedCalculation.Capacity,
          Durability: ingredientCalculation.Durability + indexedCalculation.Durability,
          Flavor: ingredientCalculation.Flavor + indexedCalculation.Flavor,
          Texture: ingredientCalculation.Texture + indexedCalculation.Texture,
          Calories: ingredientCalculation.Calories + indexedCalculation.Calories);
      }
    }
  }

  private readonly record struct Ingredient(int Capacity, int Durability, int Flavor, int Texture, int Calories);

  private readonly record struct IngredientCalculation(int Capacity, int Durability, int Flavor, int Texture, int Calories);

  [GeneratedRegex(
    @"(?<ingredient>\w+): capacity (?<capacity>-?\d+), durability (?<durability>-?\d+), flavor (?<flavor>-?\d+), texture (?<texture>-?\d+), calories (?<calories>-?\d+)")]
  private static partial Regex IngredientRegex();
}