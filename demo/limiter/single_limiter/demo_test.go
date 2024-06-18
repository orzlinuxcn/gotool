package single_limiter

import (
	"fmt"
	"testing"
)

func TestSingleLimiter(t *testing.T) {
	//ctx := context.Background()
	//var b int64 = 3
	//qpsLimiter := limiter.NewLimiter("", 3, b)
	//for i := 0; i < 30; i++ {
	//	go func() {
	//		sleepTime := int64(rand.Float64() * 1000)
	//		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	//		qpsLimiter.Wait(ctx)
	//		println(time.Now().String())
	//		println(qpsLimiter.Tokens())
	//	}()
	//}
	//time.Sleep(20 * time.Second)
	//qpsLimiter.Allow()
	//println(qpsLimiter.Tokens())
	println(fmt.Sprintf("%06d", 132))
}
