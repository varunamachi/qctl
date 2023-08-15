package param

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/varunamachi/libx/errx"
)

func bindParamVal(etx echo.Context, id string, paramType Type) (any, error) {

	var data any
	switch paramType {
	case PtConstant:
		data = ""
	case PtBoolean:
		data = false
	case PtTristate:
		data = None
	case PtChoice:
		data = ""
	case PtNumber:
		data = 0
	case PtRange:
		data = Range{Start: 0, End: 0}
	case PtDate:
		data = time.Now()
	case PtDateRange:
		data = DateRnage{Start: time.Now(), End: time.Now()}
	}

	if err := etx.Bind(&data); err != nil {
		return nil, errx.BadReqX(err, "failed to get value for param '%s'", id)
	}

	return data, nil

}
