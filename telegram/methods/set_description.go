package methods

import (
	"encoding/json"
	"github.com/NotFriendly/basebot/telegram/types"
)

// 发送机器人描述
type setDescription struct {
	Description   string `json:"description"`   // 机器人描述
	Language_code string `json:"language_code"` // 语言代码
}

func (bot *BotExt) setDescription(description string, language_code string) (*types.Message, error) {
	request := setDescription{
		Description:   description,
		Language_code: language_code,
	}

	res, err := bot.Call("sendDocument", request)
	if err != nil {
		return nil, err
	}

	resonpe := SendMessageResonpe{}
	err = json.Unmarshal(res, &resonpe)
	if err != nil {
		return nil, err
	}
	return resonpe.Result, nil
}
