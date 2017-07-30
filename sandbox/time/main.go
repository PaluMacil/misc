package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now:")
	n := time.Now()
	fmt.Println(n)
	fmt.Println(n.Location())
	fmt.Println("UTC:")
	u := n.UTC()
	fmt.Println(u)
	fmt.Println(u.Location())
	fmt.Println("Diff:")
	fmt.Println(n.Sub(u))
	b, e := u.GobEncode()
	fmt.Println(b, e)
	fmt.Println(u.Unix())
	fmt.Println("CET:")
	// IANA.org timezones (same as most OS databases): https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	loc, e := time.LoadLocation("CET")
	fmt.Println(loc, e)
	c := n.In(loc)
	fmt.Println(c)
	/* JAVASCRIPT DATES
	   (new Date()).toISOString() // Requires shim in IE8.
	   "2017-07-30T18:17:45.260Z"
	*/
	jsTime := "2017-07-30T18:17:45.260Z"
	fmt.Println("Parse JavaScript dates from (new Date()).toISOString()",
		jsTime)
	j, e := time.Parse(time.RFC3339, jsTime)
	fmt.Println(j, e)
}

/* GODOC HIGHLIGHTS

const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)

func Now
func Now() Time
Now returns the current local time.

func Parse
func Parse(layout, value string) (Time, error)
Parse parses a formatted string and returns the time value it represents. The layout defines the format by showing how
the reference time, defined to be

Mon Jan 2 15:04:05 -0700 MST 2006
would be interpreted if it were the value; it serves as an example of the input format. The same interpretation will
then be made to the input string.

Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard and convenient representations of the
reference time. For more information about the formats and the definition of the reference time, see the documentation
for ANSIC and the other constants defined by this package. Also, the executable example for time.Format demonstrates
the working of the layout string in detail and is a good reference.

Elements omitted from the value are assumed to be zero or, when zero is impossible, one, so parsing "3:04pm" returns
the time corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is 0, this time is before the zero
Time). Years must be in the range 0000..9999. The day of the week is checked for syntax but it is otherwise ignored.

In the absence of a time zone indicator, Parse returns a time in UTC.

When parsing a time with a zone offset like -0700, if the offset corresponds to a time zone used by the current
location (Local), then Parse uses that location and zone in the returned time. Otherwise it records the time as being
in a fabricated location with time fixed at the given zone offset.

No checking is done that the day of the month is within the month's valid dates; any one- or two-digit value is
accepted. For example February 31 and even February 99 are valid dates, specifying dates in March and May. This
behavior is consistent with time.Date.

When parsing a time with a zone abbreviation like MST, if the zone abbreviation has a defined offset in the current
location, then that offset is used. The zone abbreviation "UTC" is recognized as UTC regardless of location. If the
zone abbreviation is unknown, Parse records the time as being in a fabricated location with the given zone abbreviation
and a zero offset. This choice means that such a time can be parsed and reformatted with the same layout losslessly,
but the exact instant used in the representation will differ by the actual zone offset. To avoid such problems, prefer
time layouts that use a numeric zone offset, or use ParseInLocation.

Example
func ParseInLocation
func ParseInLocation(layout, value string, loc *Location) (Time, error)
ParseInLocation is like Parse but differs in two important ways. First, in the absence of time zone information, Parse
interprets a time as UTC; ParseInLocation interprets the time as in the given location. Second, when given a zone
offset or abbreviation, Parse tries to match it against the Local location; ParseInLocation uses the given location.

func (Time) Before

func (t Time) Before(u Time) bool
Before reports whether the time instant t is before u.

func (Time) Clock
func (t Time) Clock() (hour, min, sec int)
Clock returns the hour, minute, and second within the day specified by t.

func (Time) Date
func (t Time) Date() (year int, month Month, day int)
Date returns the year, month, and day in which t occurs.

func (Time) Day
func (t Time) Day() int
Day returns the day of the month specified by t.

func (Time) Equal
func (t Time) Equal(u Time) bool
Equal reports whether t and u represent the same time instant. Two times can be equal even if they are in different
locations. For example, 6:00 +0200 CEST and 4:00 UTC are Equal. Do not use == with Time values.

func (Time) Format
func (t Time) Format(layout string) string

type Duration int64
A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the
largest representable duration to approximately 290 years.

const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)

func AfterFunc
func AfterFunc(d Duration, f func()) *Timer
AfterFunc waits for the duration to elapse and then calls f in its own goroutine. It returns a Timer that can be used
to cancel the call using its Stop method.
*/
