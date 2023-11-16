package sample

import (
	"fmt"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
)

func GetBucketLocation() {
	// New client
	client := NewClient()
	ret, err := client.GetBucketLocation(bucketName)
	if err != nil {
		HandleError(err)
	}

	fmt.Println(ret.DataLocationType == oos.DataLocationTypeSpecified)
	fmt.Println(ret)
}
