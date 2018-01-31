<div align="center">
  <h1>match-sorter</h1>

  <p>Simple, expected, and deterministic best-match sorting of an array in Go Lang</p>
</div>

## The problem

1. You have a list of dozens, hundreds, or thousands of items
2. You want to filter and sort those items intelligently (maybe you have a filter input for the user)
3. You want simple, expected, and deterministic sorting of the items (no fancy math algorithm that fancily changes the sorting as they type)

## This solution

This follows a simple and sensible (user friendly) algorithm that makes it easy for you to filter and sort a list of items based on given input. Items are ranked based on sensible criteria that result in a better user experience.

To explain the ranking system, I'll use countries as an example:

1. **CASE SENSITIVE EQUALS**: Case-sensitive equality trumps all. These will be first. (ex. `France` would match `France `, but not `france`)
2. **EQUALS**: Case-insensitive equality (ex. `France` would match `france`)
3. **STARTS WITH**: If the item starts with the given value (ex. `Sou` would match `South Korea` or `South Africa`)
4. **WORD STARTS WITH**: If the item has multiple words, then if one of those words starts with the given value (ex. `Repub` would match `Dominican Republic`)
5. **CASE STARTS WITH**: If the item has a defined case (`camelCase`, `PascalCase`, `snake_case` or `kebab-case`), then if one of the parts starts with the given value (ex. `kingdom` would match `unitedKingdom` or `united_kingdom`)
6. **CASE ACRONYM** If the item's case matches the synonym (ex. `uk` would match `united-kingdom` or `UnitedKingdom`)
7. **CONTAINS**: If the item contains the given value (ex. `ham` would match `Bahamas`)
8. **ACRONYM**: If the item's acronym is the given value (ex. `us` would match `United States`)
9. **SIMPLE MATCH**: If the item has letters in the same order as the letters of the given value (ex. `iw` would match `Zimbabwe`, but not `Kuwait` because it must be in the same order). Furthermore, if the item is a closer match, it will rank higher (ex. `ua` matches `Uruguay` more closely than `United States of America`, therefore `Uruguay` will be ordered before `United States of America`)

This ranking seems to make sense in people's minds. At least it does in mine. Feedback welcome!

## LICENSE

MIT
