/**
 * @Author Mr.LiuQH
 * @Description RSA测试使用
 * @Date 2021/7/1 4:40 下午
 **/
package crypto

import (
	"52lu/go-study-example/package/crypto"
	"fmt"
	"testing"
)

// 测试生成密钥对
func TestGenerateKey(t *testing.T) {
	key, err := crypto.GenerateRSAKey(1024,"../../tmp")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v",key)
}