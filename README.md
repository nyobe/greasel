Todo:
- Preload api / nested eager loading
  - Ecto Preload: https://hexdocs.pm/ecto/Ecto.Query.html#preload/3
  - Diesel BelongsTo: https://github.com/diesel-rs/diesel/issues/89
  - Datomic Pull: https://docs.datomic.com/cloud/query/query-pull.html
  - Toucan Hydration: https://github.com/metabase/toucan/blob/master/docs/hydration.md
- Subqueries - handle by converting query to "table" with selection as fields? can't think of a typesafe way to handle this.
    ```
    sub := q.AsTable()
    q2 := From(subquery).Where(sub.IntField("boop").Eq(2))
    ```
- Partial selects - maybe structs are always fully hydrated, partial selects -> maps
- DB adapters - translate queries to eg go-pg queries
- Generate concrete field types
- Fragment escape hatch - support injecting raw sql snippets
- Function expressions
- Cast expressions
- Grouping / Aggregates

