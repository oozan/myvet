## Structs

As many, if not all, data structures are shared between multiple packages, the logical place to put them is a third separate
package such as this.

This is partially because Go makes it somewhat convenient to combine DAOs and DTOs with `db:"some" json:"other"` syntax.
However, if the discussion about DAO/DTO separation wrt null values results in separating them, it might be viable
to move the DTOs under the handlers package (as handlers use DTOs for JSON responses) and the repo package (as all the use for DAOs is there.)