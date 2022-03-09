package application

import "sync"

var BizCardServiceBean BizCardService
var once sync.Once

func SetupBizCardService() BizCardService {
	once.Do(func() {
		if BizCardServiceBean == nil {
			BizCardServiceBean = new(BizCardServiceImpl)
		}
	})
	return BizCardServiceBean
}
