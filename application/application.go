package application

import "sync"

var BizCardServiceBean BizCardService
var once sync.Once

func SetupBizCardService() BizCardService {
	if BizCardServiceBean == nil {
		once.Do(func() {
			BizCardServiceBean = new(BizCardServiceImpl)
		})
	}
	return BizCardServiceBean
}
