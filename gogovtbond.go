package gogovtbond

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// BondValue represents the value of a bond relative to the values provided by
// the treasurey value files provided to LoadValues.
type BondValue struct {
	Series         string
	RedemtionYear  int
	RedemtionMonth int
	IssueYear      int

	IssueValueJan float64
	IssueValueFeb float64
	IssueValueMar float64
	IssueValueApr float64
	IssueValueMay float64
	IssueValueJun float64
	IssueValueJul float64
	IssueValueAug float64
	IssueValueSep float64
	IssueValueOct float64
	IssueValueNov float64
	IssueValueDec float64
}

// BondValues is a slice of BondValue structs.
type BondValues []BondValue

// BondData represents the bond data source and allows obtaining informatin on bonds.
type BondData struct {
	Values BondValues
}

// getIssueVal parses the issue month value from the data file.
func getIssueVal(l string) (float64, error) {
	if l == "NO PAY" || l == "      " {
		return 0, nil
	}

	return strconv.ParseFloat(strings.Join([]string{l[:4], l[4:]}, "."), 64)
}

// NewBondValue generates a new bond value from the tearsurey value data source line.
func NewBondValue(l string) (BondValue, error) {
	var err error

	b := BondValue{}
	b.Series = l[:1]

	b.RedemtionYear, err = strconv.Atoi(l[1:5])
	if err != nil {
		return b, err
	}

	b.RedemtionMonth, err = strconv.Atoi(l[5:7])
	if err != nil {
		return b, err
	}

	b.IssueYear, err = strconv.Atoi(l[7:11])
	if err != nil {
		return b, err
	}

	b.IssueValueJan, err = getIssueVal(l[11:17])
	if err != nil {
		return b, err
	}

	b.IssueValueFeb, err = getIssueVal(l[17:23])
	if err != nil {
		return b, err
	}

	b.IssueValueMar, err = getIssueVal(l[23:29])
	if err != nil {
		return b, err
	}

	b.IssueValueApr, err = getIssueVal(l[29:35])
	if err != nil {
		return b, err
	}

	b.IssueValueMay, err = getIssueVal(l[35:41])
	if err != nil {
		return b, err
	}

	b.IssueValueJun, err = getIssueVal(l[41:47])
	if err != nil {
		return b, err
	}

	b.IssueValueJul, err = getIssueVal(l[47:53])
	if err != nil {
		return b, err
	}

	b.IssueValueAug, err = getIssueVal(l[53:59])
	if err != nil {
		return b, err
	}

	b.IssueValueSep, err = getIssueVal(l[59:65])
	if err != nil {
		return b, err
	}

	b.IssueValueOct, err = getIssueVal(l[65:71])
	if err != nil {
		return b, err
	}

	b.IssueValueNov, err = getIssueVal(l[71:77])
	if err != nil {
		return b, err
	}

	b.IssueValueDec, err = getIssueVal(l[77:83])
	if err != nil {
		return b, err
	}

	return b, nil

}

// Load loads the bond values from an io.Reader interface. The values must be returned
// in the same format provided by the US Treasury.
// https://www.treasurydirect.gov/indiv/tools/tools_savingsbondvalues_specifications.htm
func (b *BondData) Load(f io.Reader) error {
	r := bufio.NewReader(f)

	// Read until we reach the end.
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			break
		}

		nb, err := NewBondValue(l)
		if err != nil {
			return err
		}

		b.Values = append(b.Values, nb)
	}

	return nil
}

// LoadFromFile loads the bond data from the file specified by path.
func (b *BondData) LoadFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return b.Load(f)
}

// BondValue returns the bond value reported by the data file loaded.
func (b *BondData) BondValue(series string, redemtionYear int, redemtionMonth int, issueYear int, issueMonth int, value int) float64 {

	for _, b := range b.Values {

		if b.Series == series && b.RedemtionYear == redemtionYear && b.RedemtionMonth == redemtionMonth && b.IssueYear == issueYear {
			switch issueMonth {
			case 1:
				return b.IssueValueJan * (float64(value) / 25)
			case 2:
				return b.IssueValueFeb * (float64(value) / 25)
			case 3:
				return b.IssueValueMar * (float64(value) / 25)
			case 4:
				return b.IssueValueApr * (float64(value) / 25)
			case 5:
				return b.IssueValueMay * (float64(value) / 25)
			case 6:
				return b.IssueValueJun * (float64(value) / 25)
			case 7:
				return b.IssueValueJul * (float64(value) / 25)
			case 8:
				return b.IssueValueAug * (float64(value) / 25)
			case 9:
				return b.IssueValueSep * (float64(value) / 25)
			case 10:
				return b.IssueValueOct * (float64(value) / 25)
			case 11:
				return b.IssueValueNov * (float64(value) / 25)
			case 12:
				return b.IssueValueDec * (float64(value) / 25)
			}
		}
	}

	return 0
}
