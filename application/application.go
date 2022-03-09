package application

import "sync"

var BizCardService2 BizCardService
var once sync.Once

func SetupBizCardService() BizCardService {
	once.Do(func() {
		if BizCardService2 == nil {
			BizCardService2 = new(BizCardServiceImpl)
		}
	})
	return BizCardService2
}
