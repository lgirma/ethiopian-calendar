# ethiopian-calendar

Ethiopian calendar library for Go

## Installation

Install using `go get` as:

```shell
go get github.com/lgirma/ethiopian-calendar
```

## Usage

### Date Conversions

To convert a date in Ethiopian calendar to Gregorian calendar, use `ToGregorian()` as:

```go
// Convert 8/10/2014 from EC to GC
ethcal.ToGregorian(ethcal.NewDate(2014, 10, 8))
```

To convert a date in Gregorian calendar to Ethiopian calendar, use `ToEthiopian()` as:

```go
// Convert 01/01/2022 from GC to EC
ethcal.ToEthiopian(ethcal.NewDate(2022, 01, 01))
```

## Moving Feasts

The Easter and related feasts in the Ethiopian Orthodox Tewahedo Church are not pinned to specific dates.

To get on which day of year a feast will be at:

```go
// Get date of start of the Great lent for the year 2014 in Ethiopian Calendar
ethcal.GetFeast(ethcal.FEAST_ABIY_TSOM, 2014)

// Get date of Easter for the year 2013 in Ethiopian Calendar
ethcal.GetFeast(ethcal.FEAST_TINSAE, 2013)
```

The following feasts are available:

```
FEAST_ABIY_TSOM = 14            // Or can use FEAST_GREAT_LENT
FEAST_DEBREZEIT = 41
FEAST_HOSAENA = 62
FEAST_SIQLET = 67               // Or can use FEAST_GOOD_FRIDAY
FEAST_TINSAE = 69               // Or can use FEAST_EASTER
FEAST_RIKBE_KAHINAT = 93
FEAST_ERGET = 108               // Or can use FEAST_ASCENSION
FEAST_PERAQLITOS = 118          // Or can use FEAST_PENTECOST
FEAST_TSOME_HAWARIAT = 119      // Or can use FEAST_APOSTLES_FAST
```