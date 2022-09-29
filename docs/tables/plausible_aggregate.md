# Table: plausible_aggregate

Aggregate a number of metrics over a certain time period for a given domain.

## Examples

### Get number of unique visitors for the last 30 day

```sql
select
  visitors
from
  plausible_aggregate
where
  domain='example.net'
```

### Get number of unique visitors for today

```sql
select
  visitors
from
  plausible_aggregate
where
  domain='example.net'
  and period='day'
```

### Get number of unique visitors for a specific date

```sql
select
  visitors
from
  plausible_aggregate
where
  domain='example.net'
  and period='custom'
  and date='2022-09-15,2022-09-15'
```

### Get number of unique visitors for the last 30 day with the variation in % from the previous period

```sql
select
  visitors,
  visitors_change
from
  plausible_aggregate
where
  domain='example.net'
```
