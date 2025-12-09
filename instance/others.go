package instance

import (
	"fmt"
	"strings"

	"github.com/BridgeSenseDev/Dank-Memer-Grinder/gateway"
	"github.com/BridgeSenseDev/Dank-Memer-Grinder/utils"
)

func (in *Instance) Others(message gateway.EventMessage) {
	embed := in.FetchEmbed(message, 0)
	if embed.Title == "You have an unread alert!" && in.Cfg.ReadAlerts && !in.IsPaused() {
		err := in.SendCommand("alert", nil, true)

		if err != nil {
			utils.Log(utils.Others, utils.Error, in.SafeGetUsername(), fmt.Sprintf("Failed to send /alert command: %s", err.Error()))
		}
	}

	if strings.Contains(strings.ToLower(embed.Title), "maintenance") {
		utils.Log(utils.Important, utils.Info, in.SafeGetUsername(), "Global toggle has been switched due to a Dank Memer maintenance. Check if the update is safe before continuing to grind")
		in.Cfg.State = false
		in.UpdateConfig(in.Cfg)
		utils.EmitEventIfNotCLI("configUpdate", in.Cfg)
	}
}
