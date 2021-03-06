package ethiopian_calendar

// DayOfYear Represents a date structure with Day and Month
type DayOfYear struct {
	Day int
	Month int
}

// NewDay Creates a new DayOfYear structure
func NewDay(month int, day int) DayOfYear {
	return DayOfYear{Month: month, Day: day}
}

// NewDate Creates a new Date structure
func NewDate(year int, month int, day int) Date {
	return Date{Year: year, Month: month, Day: day}
}

// Date Represents a date structure with Day, Month and Year
type Date struct {
	Day int
	Month int
	Year int
}

const EVANGELIST_MATTHEW = 1
const EVANGELIST_MARK = 2
const EVANGELIST_LUKE = 3
const EVANGELIST_JOHN = 4

const MONTH_MESKEREM = 1
const MONTH_TIKIMT = 2
const MONTH_HIDAR = 3
const MONTH_TAHISAS = 4
const MONTH_TIR = 5
const MONTH_YEKATIT = 6
const MONTH_MEGABIT = 7
const MONTH_MIYAZIYA = 8
const MONTH_GINBOT = 9
const MONTH_SENE = 10
const MONTH_HAMLE = 11
const MONTH_NEHASE = 12
const MONTH_PAGUMEN = 13

const MONTH_SEP = 1
const MONTH_OCT = 2
const MONTH_NOV = 3
const MONTH_DEC = 4
const MONTH_JAN = 5
const MONTH_FEB = 6
const MONTH_MAR = 7
const MONTH_APR = 8
const MONTH_MAY = 9
const MONTH_JUN = 10
const MONTH_JUL = 11
const MONTH_AUG = 12

const DAY_MON = 1
const DAY_TUE = 2
const DAY_WED = 3
const DAY_THU = 4
const DAY_FRI = 5
const DAY_SAT = 6
const DAY_SUN = 7

const FEAST_ABIY_TSOM = 14
const FEAST_GREAT_LENT = 14
const FEAST_DEBREZEIT = 41
const FEAST_HOSAENA = 62
const FEAST_SIQLET = 67
const FEAST_GOOD_FRIDAY = 67
const FEAST_TINSAE = 69
const FEAST_EASTER = 69
const FEAST_RIKBE_KAHINAT = 93
const FEAST_ERGET = 108
const FEAST_ASCENSION = 108
const FEAST_PERAQLITOS = 118
const FEAST_PENTECOST = 118
const FEAST_TSOME_HAWARIAT = 119
const FEAST_APOSTLES_FAST = 119