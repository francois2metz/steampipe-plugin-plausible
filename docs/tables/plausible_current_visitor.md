# Table: plausible_current_visitor

Returns the number of live visitors for a given domain.

## Examples

### Live visitors of the domain

```sql
select
  count
from
  plausible_current_visitor
where
  domain='example.net'
```
