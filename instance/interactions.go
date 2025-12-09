package instance

import (
	"fmt"
	"strings"

	"github.com/BridgeSenseDev/Dank-Memer-Grinder/discord/types"
	"github.com/BridgeSenseDev/Dank-Memer-Grinder/gateway"
	"github.com/BridgeSenseDev/Dank-Memer-Grinder/utils"
)

func (in *Instance) SendChatMessage(content string, delay bool) error {
	if !in.Cfg.State || !in.AccountCfg.State {
		return nil
	}

	if delay {
		<-utils.Sleep(utils.RandSeconds(in.Cfg.Cooldowns.CommandInterval.MinSeconds, in.Cfg.Cooldowns.CommandInterval.MaxSeconds))
	}

	return in.Client.SendChatMessage(content)
}

func (in *Instance) SendCommand(name string, options map[string]string, delay bool) error {
	if !in.Cfg.State || !in.AccountCfg.State {
		return nil
	}

	if delay {
		<-utils.Sleep(utils.RandSeconds(in.Cfg.Cooldowns.CommandInterval.MinSeconds, in.Cfg.Cooldowns.CommandInterval.MaxSeconds))
	}

	return in.Client.SendCommand(name, options)
}

func (in *Instance) SendSubCommand(name string, subCommandName string, options map[string]string, delay bool) error {
	if !in.Cfg.State || !in.AccountCfg.State {
		return nil
	}

	if delay {
		<-utils.Sleep(utils.RandSeconds(in.Cfg.Cooldowns.CommandInterval.MinSeconds, in.Cfg.Cooldowns.CommandInterval.MaxSeconds))
	}

	return in.Client.SendSubCommand(name, subCommandName, options)
}

func (in *Instance) FetchEmbed(message gateway.EventMessage, index int) types.Embed {
	fmt.Println("we here")
	if len(message.Embeds) != 0 {
		return message.Embeds[index]
	} else {
		fmt.Println("we here2")
		return types.Embed{}
	}

}

func (in *Instance) FindComponentContent(message types.MessageComponent, itemToSearch string) (bool, types.MessageComponent) {
	search := strings.ToLower(itemToSearch)

	switch item := message.(type) {
	case *types.ActionsRow:
		for _, child := range item.Components {
			match, res := in.FindComponentContent(child, itemToSearch)
			if match {
				return true, res
			}
		}

	case *types.Container:
		for _, child := range item.Components {
			match, res := in.FindComponentContent(child, itemToSearch)
			if match {
				return true, res
			}
		}

	case *types.Section:
		for _, child := range item.Components {
			match, res := in.FindComponentContent(child, itemToSearch)
			if match {
				return true, res
			}
		}

	case *types.Button:
		if strings.Contains(strings.ToLower(item.Label), search) ||
			strings.Contains(strings.ToLower(item.CustomID), search) {
			return true, item
		}

	case *types.SelectMenu:
		if strings.Contains(strings.ToLower(item.CustomID), search) {
			return true, item
		}

	case *types.TextInput:
		if strings.Contains(strings.ToLower(item.CustomID), search) ||
			strings.Contains(strings.ToLower(item.Value), search) {
			return true, item
		}

	case *types.TextDisplay:
		if strings.Contains(strings.ToLower(item.Content), search) {
			return true, item
		}

	case *types.Label:
		if strings.Contains(strings.ToLower(item.Description), search) {
			return true, item
		}
	}

	return false, nil
}

func (in *Instance) ClickButton(message gateway.EventMessage, row int, column int) error {
	if !in.Cfg.State || !in.AccountCfg.State {
		return nil
	}

	<-utils.Sleep(utils.RandSeconds(in.Cfg.Cooldowns.ButtonClickDelay.MinSeconds, in.Cfg.Cooldowns.ButtonClickDelay.MaxSeconds))

	err := in.Client.ClickButton(message, row, column)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (in *Instance) ClickDmButton(message gateway.EventMessage, row int, column int) error {
	if !in.Cfg.State || !in.AccountCfg.State {
		return nil
	}

	<-utils.Sleep(utils.RandSeconds(in.Cfg.Cooldowns.ButtonClickDelay.MinSeconds, in.Cfg.Cooldowns.ButtonClickDelay.MaxSeconds))
	return in.Client.ClickDmButton(message, row, column)
}

func (in *Instance) ChooseSelectMenu(message gateway.EventMessage, row int, column int, values []string) error {
	if !in.Cfg.State || !in.AccountCfg.State {
		return nil
	}

	<-utils.Sleep(utils.RandSeconds(in.Cfg.Cooldowns.ButtonClickDelay.MinSeconds, in.Cfg.Cooldowns.ButtonClickDelay.MaxSeconds))
	return in.Client.ChooseSelectMenu(message, row, column, values)
}

func (in *Instance) SubmitModal(modal gateway.EventModalCreate) error {
	if !in.Cfg.State || !in.AccountCfg.State {
		return nil
	}

	<-utils.Sleep(utils.RandSeconds(in.Cfg.Cooldowns.ButtonClickDelay.MinSeconds, in.Cfg.Cooldowns.ButtonClickDelay.MaxSeconds))
	return in.Client.SubmitModal(modal)
}
