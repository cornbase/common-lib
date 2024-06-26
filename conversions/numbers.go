package conversions

import (
	"math/big"

	"github.com/jackc/pgtype"
)

func NumericToRat(n pgtype.Numeric) (*big.Rat, error) {
	rat := big.NewRat(1, 1)
	err := n.AssignTo(rat)
	if err != nil {
		return nil, err
	}
	return rat, nil
}

func NumericToFloat(n pgtype.Numeric) (*big.Float, error) {
	rat, err := NumericToRat(n)
	if err != nil {
		return nil, err
	}
	return big.NewFloat(0).SetRat(rat), nil
}

func NumericToInt(n pgtype.Numeric) (*big.Int, error) {
	f, err := NumericToFloat(n)
	if err != nil {
		return nil, err
	}
	out, _ := f.Int(nil)
	return out, nil
}

func MustNumericToInt(n pgtype.Numeric) *big.Int {
	out, err := NumericToInt(n)
	if err != nil {
		panic(err)
	}
	return out
}

func FloatToNumeric(f *big.Float) (pgtype.Numeric, error) {
	n := &pgtype.Numeric{}
	err := n.Set(f.String())
	return *n, err
}

func MustFloatToNumeric(f *big.Float) pgtype.Numeric {
	n, err := FloatToNumeric(f)
	if err != nil {
		panic(err)
	}
	return n
}

func RatToNumeric(r *big.Rat) (pgtype.Numeric, error) {
	return FloatToNumeric(big.NewFloat(0).SetRat(r))
}

func IntToNumeric(i *big.Int) (pgtype.Numeric, error) {
	n := &pgtype.Numeric{}
	err := n.Set(i.String())
	return *n, err
}

func MustIntToNumeric(i *big.Int) pgtype.Numeric {
	n, err := IntToNumeric(i)
	if err != nil {
		panic(err)
	}
	return n
}
