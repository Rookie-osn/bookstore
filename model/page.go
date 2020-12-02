package model

type Page struct {
	Books       []*Book //存储页面内的book变量的切片
	PageSize    int64   //每页图书数量
	PageNo      int64   //当前页
	TotalPageNo int64   //总页数，通过计算得到
	TotalRecord int64   //总图书量，从数据库获取
	MinPrice    string  //过滤的最低价格
	MaxPrice    string  //过滤的最高价格
	IsLogin     bool    //是否已登录
	Username    string  //用户名
}

//判断是否有上一页
func (p *Page) IsHasPrev() bool {
	if p.PageNo > 1 {
		return true
	}
	return false
}

//判断是否有下一页
func (p *Page) IsHasNext() bool {
	if p.PageNo < p.TotalPageNo {
		return true
	}
	return false
}

//获取上一页页码
func (p *Page) GetPrevPageNo() int64 {
	if p.PageNo > 1 {
		return p.PageNo - 1
	}
	return 1
}

//获取下一页页码
func (p *Page) GetNextPageNo() int64 {
	if p.PageNo < p.TotalPageNo {
		return p.PageNo + 1
	}
	return p.TotalPageNo
}