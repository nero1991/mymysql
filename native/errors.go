package native

import (
	"errors"
)

var (
	SEQ_ERROR             = errors.New("packet sequence error")
	PKT_ERROR             = errors.New("malformed packet")
	PKT_LONG_ERROR        = errors.New("packet too long")
	UNEXP_NULL_LCS_ERROR  = errors.New("unexpected null LCS")
	UNEXP_NULL_LCB_ERROR  = errors.New("unexpected null LCB")
	UNEXP_NULL_DATE_ERROR = errors.New("unexpected null datetime")
	UNEXP_NULL_TIME_ERROR = errors.New("unexpected null time")
	UNK_RESULT_PKT_ERROR  = errors.New("unexpected or unknown result packet")
	NOT_CONN_ERROR        = errors.New("not connected")
	ALREDY_CONN_ERROR     = errors.New("not connected")
	BAD_RESULT_ERROR      = errors.New("unexpected result")
	UNREADED_ROWS_ERROR   = errors.New("there are unreaded rows")
	BIND_COUNT_ERROR      = errors.New("wrong number of values for bind")
	BIND_UNK_TYPE         = errors.New("unknown value type for bind")
	RESULT_COUNT_ERROR    = errors.New("wrong number of result columns")
	BAD_COMMAND_ERROR     = errors.New("comand isn't text SQL nor *Stmt")
	WRONG_DATE_LEN_ERROR  = errors.New("wrong datetime/timestamp length")
	WRONG_TIME_LEN_ERROR  = errors.New("wrong time length")
	UNK_MYSQL_TYPE_ERROR  = errors.New("unknown MySQL type")
	WRONG_PARAM_NUM_ERROR = errors.New("wrong parameter number")
	UNK_DATA_TYPE_ERROR   = errors.New("unknown data source type")
	SMALL_PKT_SIZE_ERROR  = errors.New("specified packet size is to small")
)