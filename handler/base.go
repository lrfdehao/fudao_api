package handler

type Res struct {
	Retcode int         `json:"retcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func res(code int, desc string, data interface{}) Res {
	var res Res
	res.Retcode = code
	res.Msg = desc
	res.Data = data
	return res
}
