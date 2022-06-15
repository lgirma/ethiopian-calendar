package ethiopian_calendar

type DayOfYear struct {
	Day int
	Month int
}

func NewDay(month int, day int) DayOfYear {
	return DayOfYear{Month: month, Day: day}
}

func NewDate(year int, month int, day int) Date {
	return Date{Year: year, Month: month, Day: day}
}

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

const DAY_MON = 1
const DAY_TUE = 2
const DAY_WED = 3
const DAY_THU = 4
const DAY_FRI = 5
const DAY_SAT = 6
const DAY_SUN = 7

const FEAST_ABIY_TSOM = 14
const FEAST_DEBREZEIT = 41
const FEAST_HOSAENA = 62
const FEAST_SIQLET = 67
const FEAST_TINSAE = 69
const FEAST_RIKBE_KAHINAT = 93
const FEAST_ERGET = 108
const FEAST_PERAQLITOS = 118
const FEAST_TSOME_HAWARIAT = 119