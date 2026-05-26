# Filters and piped stdin

GoWhen can act as a shell filter. When stdin is piped, each non-empty input line is automatically parsed as the working date before the command chain runs.

```sh
cat dates.txt | GoWhen format 2006-01-02
```

This is equivalent to running the command chain once for each line, with a hidden parse step first:

```text
parse . <stdin-line> format 2006-01-02
```

Commands still keep their normal argument rules. Piped stdin supplies the working date; it does not fill missing command arguments.

For example, this compares each input date against a fixed command-line date:

```sh
cat dates.txt | GoWhen diff . 2026-06-01
```

But this is still invalid, because `diff` requires both arguments:

```sh
cat dates.txt | GoWhen diff .
```

## keep and drop

`keep` and `drop` filter output based on the current working date.

```sh
GoWhen keep <selector>
GoWhen drop <selector>
```

`keep` prints only dates matching the selector.

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d keep monday
```

`drop` suppresses dates matching the selector.

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d drop weekend
```

They also work with piped input:

```sh
printf '%s\n' 2026-05-30 2026-05-31 2026-06-01 | GoWhen drop weekend format 2006-01-02
```

Output:

```text
2026-06-01
```

## Selectors

Current selectors are:

```text
weekday
weekend

mon | monday
tue | tuesday
wed | wednesday
thu | thursday
fri | friday
sat | saturday
sun | sunday
```

## Examples

Keep weekdays from piped input:

```sh
cat dates.txt | GoWhen keep weekday format 2006-01-02
```

Drop weekends from piped input:

```sh
cat dates.txt | GoWhen drop weekend format 2006-01-02
```

Keep Mondays from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d keep mon
```

Drop Saturdays from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d drop saturday
```

Keep today only if today is a weekday:

```sh
GoWhen keep weekday
```

Drop today if today is a weekend:

```sh
GoWhen drop weekend
```
