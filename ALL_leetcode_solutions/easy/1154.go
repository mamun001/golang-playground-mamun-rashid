// easy 1154, ELO 1199
// Day of the year


import "strconv"

func month_days(s string) int {

    ans := 0

    switch s {
        case "01":
           ans = 0
        case "02":
           ans = 31
        case "03":
           ans = 59
        case "04":
           ans = 90
        case "05":
           ans = 120
        case "06":
           ans = 151
        case "07":
           ans = 181
        case "08":
           ans = 212
        case "09":
           ans = 243
        case "10":
           ans = 273
        case "11":
           ans = 304
        case "12":
           ans = 334
    }

    return ans
}

func dayOfYear(date string) int {

    year := date[0:4]
    //fmt.Println(year)
    year_int,_ := strconv.Atoi(year)
    //fmt.Println(year_int)


    month := date[5:7]
    month_int,_ := strconv.Atoi(month)
    //fmt.Println("month_int",month_int)

    day := date[8:10]
    if day[0] == '0' {
        day = date[9:10]
    }
    day_int,_ := strconv.Atoi(day)
    //fmt.Println("day_int",day_int)


    mo_day_days :=month_days(month)+day_int
    if year_int % 4 == 0 && year_int != 1900 && month_int > 2 {
        mo_day_days = mo_day_days+1
    }
    //fmt.Println("mo_day_days",mo_day_days)


    return mo_day_days
    
}
