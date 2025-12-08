package gateway

import (
	"fmt"
	"regexp"
	"time"

	"github.com/BridgeSenseDev/Dank-Memer-Grinder/utils"
	"github.com/valyala/fasthttp"
)

const fallbackBuildNumber = "9999"

var requestClient = fasthttp.Client{
	ReadBufferSize:                8192,
	ReadTimeout:                   time.Second * 5,
	WriteTimeout:                  time.Second * 5,
	NoDefaultUserAgentHeader:      true,
	DisableHeaderNamesNormalizing: true,
	DisablePathNormalizing:        true,
}

func (g *gatewayImpl) getLatestBuild() string {
	sentery_asset_regex, _ := regexp.Compile(`assets/(sentry\.\w+)\.js`)
	build_num_regex, _ := regexp.Compile(`buildNumber\D+(\d+)"`)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI("https://discord.com/login")

	// Apply headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://discord.com/login")
	req.Header.Set("Sec-Ch-Ua", `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Dest", "script")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/124.0.0.0 Safari/537.36",
	)

	if err := requestClient.Do(req, resp); err != nil {
		utils.Log(utils.Discord, utils.Error, g.SafeGetUsername(), fmt.Sprintf("Error fetching build number: %v", err))
		return fallbackBuildNumber
	} else {
		body := string(resp.Body())
		match := sentery_asset_regex.FindStringSubmatch(body)
		if match != nil {
			sentry_url := fmt.Sprintf("https://static.discord.com/assets/%v.js", match[1])
			req.SetRequestURI(sentry_url)
			if err := requestClient.Do(req, resp); err != nil {
				return fallbackBuildNumber
			} else {
				body := string(resp.Body())
				match := build_num_regex.FindStringSubmatch(body)
				return string(match[1])
			}

		} else {
			return fallbackBuildNumber
		}
	}

}

func (g *gatewayImpl) mustGetLatestBuild() string {
	build := g.getLatestBuild()
	if build == fallbackBuildNumber {
		panic("Failed to get the latest build number")
	}
	return build
}
