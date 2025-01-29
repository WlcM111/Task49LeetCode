# Group Anagrams

This program groups anagrams from a given list of strings. Anagrams are words or phrases formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Function

- **groupAnagrams(strs []string) [][]string**: Groups anagrams together from the input list of strings `strs`.

## How It Works

1. **Initial Checks**: If the input list is empty, it returns an empty result. If the list contains only one string, it returns a list containing that single string.
2. **Anagram Grouping**:
   - Iterates through each string in the input list.
   - Converts the string into a sorted version of its characters (using `sort.Slice`).
   - Uses the sorted string as a key in a map to group anagrams together.
3. **Result Construction**: After processing all strings, it collects the grouped anagrams from the map and returns them as a list of lists.

## Example

For the input `["eat", "tea", "tan", "ate", "nat", "bat"]`, the program will return: `[["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]`.
