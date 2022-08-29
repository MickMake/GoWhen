package cmd

/*
Examples:
parse date "Sat 01 Jul 1967 09:42:42 AEST" add "20d" format "2006-01-02T15:04:05"
add -- '-1y 12M -1w +7d -2h 120m -2s +2000ms' format '2006-01-02 15:04:05'
tz "UTC" format '2006-01-02 15:04:05'

*/

const (
    True = "YES"
    False = "NO"
)

const (
    Examples = `
### Parsing.
Parse the date string "Sat 01 Jul 1967 09:42:42 AEST".
    % GoWhen parse "Sat 31 Jul 1967 09:42:42 AEST" ""
    1967-07-31T09:42:42+10:00

Parse the date string "1967-07-01 09:42:42" with custom format "2006-01-02 15:04:05".
    % GoWhen parse "1967-07-01 09:42:42" "2006-01-02 15:04:05"
    1967-07-01T09:42:42Z

Print UNIX epoch
    % GoWhen parse "epoch" ""
    1970-01-01T00:00:00Z


### Adding
Parse today's date and add "20 days".
    % GoWhen parse "today" "" add "20d"
    2022-09-18T06:35:07+10:00

Parse today's date and add "2000 microseconds".
    % GoWhen parse "now" "" add "2000us"

Parse today's date and add "-1 year, +12 months, -1 week, +7 days, -2 hours, +120 minutes, -2 seconds, +2000 mS".
    % GoWhen add -- "-1y 12M -1w +7d -2h 120m -2s +2000ms"
    2022-08-29T06:38:36.73375+10:00


### Timezones
Convert "1967-07-01 09:42:42" to timezone "Australia/Sydney".
    % GoWhen parse "1967-07-01 09:42:42" "" timezone "Australia/Sydney"
    1967-07-01T19:42:42+10:00

Convert "1967-07-01 09:42:42" to timezone "UTC".
    % GoWhen parse "1967-07-01 09:42:42" "" timezone "UTC"
    1967-07-01T09:42:42Z

Convert "1967-07-01 09:42:42" to timezone "America/Chicago".
    % GoWhen parse "1967-07-01 09:42:42" "" timezone "America/Chicago"
    1967-07-01T04:42:42-05:00

Convert "1967-07-01 09:42:42" to timezone "Iceland".
    % GoWhen parse "1967-07-01 09:42:42" "" timezone "Iceland"
    1967-07-01T09:42:42Z


### Rounding
Round "1967-07-01 09:42:42" down to the nearest "5 minutes".
    % GoWhen parse "1967-07-01 09:42:42" "" round down 5m
    1967-07-01T09:40:00Z

Round "1967-07-01 09:42:42" up to the nearest "1 hour".
    % GoWhen parse "1967-07-01 09:42:42" "" round down 1h
    1967-07-01T10:00:00Z


### Differences
Show difference between "tomorrow" and "yesterday".
    % GoWhen parse "yesterday" . diff "tomorrow" .
    2d

Show difference between "tomorrow" and "yesterday".
    % GoWhen parse "last-week" . diff "today" .
    7d

Show difference between "now" and "2022-02-01 00:00:00".
    % GoWhen parse "now" . diff "2022-02-01 00:00:00" .
    6M 28d 7h 7m 55s

Show difference between "2022-02-01 00:00:00" and "now".
    % GoWhen parse "2020-01-01 00:00:00" . diff now .
    2y 7M 28d 8h 19m 36s


### Formatting
Format "1967-07-01 09:42:42" as "Mon Jan _2 15:04:05 MST 2006".
    % GoWhen parse "1967-07-01 09:42:42" . format UnixDate
    Sat Jul  1 09:42:42 UTC 1967

Format "1967-07-01 09:42:42" as "2006-01-02 15:04:05".
    % GoWhen parse "1967-07-01 09:42:42" . format "2006-01-02 15:04:05"
    1967-07-01 09:42:42

Print current date/time as UNIX epoch, (in seconds).
    % GoWhen parse now . format epoch
    1661754986

Print today's date as a week number.
    % parse today "" format week
    35


### Stacking
Parse the date "Sat 31 Jul 1967 09:42:42 AEST", add "20 days" and print as format "20060102/20060102_150405-webcam.jpg".
    % GoWhen parse "Sat 31 Jul 1967 09:42:42 AEST" ""  add "20d"  format "20060102/20060102_150405-webcam.jpg"
    19670820/19670820_094242-webcam.jpg

Parse the date "Sat 31 Jul 1967 09:42:42 AEST", convert to timezone "Iceland", add "1 day", round down to every "5 minutes" and print as format "2006-01-02 15:04:05".
    % GoWhen parse "1967-07-01 09:42:42" ""  timezone Iceland  add 1d  round down 5m  format "2006-01-02 15:04:05"
    1967-07-02 09:40:00

Is the date "1967-07-07 09:42:42" after "1967-07-01 09:42:42"
    % GoWhen parse "1967-07-01 09:42:42" ""  is after "1967-07-07 09:42:42" ""
    NO

Is the date "1967-07-01 09:42:42" before "1967-07-07 09:42:42"
    % GoWhen parse "1967-07-01 09:42:42" ""  is before "1967-07-07 09:42:42" ""
    YES
`
)
