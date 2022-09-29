# Table: plausible_timeseries

Return the timeseries date of the metrics for the given domain.

## Examples

### Get the number of visitors per day for the last 30 day

```sql
select
  time,
  visitors
from
  plausible_timeseries
where
  domain='example.net'
```

### Get the number of visitors per month for the last 6 months

```sql
select
  visitors
from
  plausible_timeseries
where
  domain='example.net'
  and period='6mo'
  and interval='month'
```

### ### Get the number of visitors per day for a specific period

```sql
select
  visitors
from
  plausible_timeseries
where
  domain='example.net'
  and period='custom'
  and date='2022-09-15,2022-09-20'
```
