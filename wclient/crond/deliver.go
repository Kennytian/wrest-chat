package crond

import (
	"strings"
	"time"

	"github.com/opentdp/wechat-rest/wclient"
)

func MsgDeliver(deliver, content string) {

	content = strings.TrimSpace(content)
	delivers := strings.Split(deliver, "\n")

	for _, dr := range delivers {
		logger.Warn("cron:deliver "+dr, "content", content)
		// 解析参数
		args := strings.Split(strings.TrimSpace(dr), ",")
		if len(args) < 2 {
			return
		}
		// 分渠道投递
		switch args[0] {
		case "wechat":
			time.Sleep(1 * time.Second)
			wechatMessage(args[1:], content)
		}
	}

}

// 将执行结果投递到微信

func wechatMessage(args []string, message string) int32 {

	roomid := args[0]

	wxid := ""
	if len(args) > 1 {
		wxid = args[1]
	}

	wc := wclient.Register()
	if wc == nil {
		logger.Error("cron:deliver", "error", "wclient is nil")
		return -1
	}

	return wclient.SendFlexMsg(message, wxid, roomid)

}
