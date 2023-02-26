package words_filter

import "github.com/cloudwego/kitex/pkg/klog"

func WordsFilter(s string) string {
	result, err := filterManage.Filter().Replace(s, '*')
	if err != nil {
		klog.Error(err)
		return s
	}
	return result
}
