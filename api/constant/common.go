package constant

// TZ *
const TZ = "Asia/Tokyo"

// ZapDateTimelayout *
const ZapDateTimelayout = "2006-01-02T15:04:05+09:00"

// DateTimeLayout .
const DateTimeLayout = "2006-01-02 15:04:05"

// DateTimeLayoutUntilDay .
const DateTimeLayoutUntilDay = "2006-01-02"

// TimeIndicatorJst .
const TimeIndicatorJst = "(JST)"

// TimeIndicatorMst .
const TimeIndicatorMst = "(MST)"

// JstDateTimeLayout .
const JstDateTimeLayout = DateTimeLayout + TimeIndicatorMst

// MstDateTimeLayout .
const MstDateTimeLayout = DateTimeLayout + TimeIndicatorMst
