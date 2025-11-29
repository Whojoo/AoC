using System.Runtime.InteropServices;
using System.Text.RegularExpressions;

namespace Shared;

public static class SharedExtensions
{
  extension<TKey, TValue> (Dictionary<TKey, TValue> dictionary) where TKey : notnull
  {
    public TValue GetOrAdd(TKey key, TValue value)
    {
      ref var val = ref CollectionsMarshal.GetValueRefOrAddDefault(dictionary, key, out var exists);
      if (exists)
        return val!;

      val = value;
      return value;
    }
  }

  extension(Group group)
  {
    public int IntValue() => int.Parse(group.Value);
  }
}