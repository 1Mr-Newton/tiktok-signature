package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/playwright-community/playwright-go"
)

func routeHandler(route playwright.Route) {
	if route.Request().ResourceType() == "script" {
		route.Abort()
	} else {
		route.Continue()
	}
}

func getRandomInt(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	minVal := min(a, b)
	maxVal := max(a, b)
	diff := maxVal - minVal + 1
	return minVal + rand.Intn(diff)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	app := fiber.New()

	app.Get("/sign", func(c *fiber.Ctx) error {
		urlParam := c.Query("url")
		if urlParam == "" {
			return c.Status(400).SendString("Missing URL parameter")
		}

		bogusScript, err := readFile("cmd/scripts/bogus.txt")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		navigatorScript, err := readFile("cmd/scripts/navigator.txt")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		signatureFunctionsScript, err := readFile("cmd/scripts/signature_functions.txt")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		signerScript, err := readFile("cmd/scripts/signer.txt")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		webmssdkScript, err := readFile("cmd/scripts/webmssdk.txt")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		scripts := []string{
			signerScript,
			signatureFunctionsScript,
			bogusScript,
			webmssdkScript,
			navigatorScript,
		}

		options := playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(true),
			Args: []string{
				"--disable-blink-features",
				"--disable-blink-features=AutomationControlled",
				"--disable-infobars",
				"--window-size=1920,1080",
				"--start-maximized",
				"--user-agent='Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56'",
			},
		}

		pw, err := playwright.Run()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		browser, err := pw.Chromium.Launch(options)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		emulateTemplate := pw.Devices["iPhone 11"]
		emulateTemplate.DeviceScaleFactor = float64(getRandomInt(1, 3))
		emulateTemplate.IsMobile = rand.Float64() > 0.5
		emulateTemplate.HasTouch = rand.Float64() > 0.5
		emulateTemplate.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56"
		emulateTemplate.Viewport.Width = getRandomInt(320, 1920)
		emulateTemplate.Viewport.Height = getRandomInt(320, 1920)

		newContextOptions := playwright.BrowserNewContextOptions{
			BypassCSP:         playwright.Bool(true),
			DeviceScaleFactor: &emulateTemplate.DeviceScaleFactor,
			IsMobile:          &emulateTemplate.IsMobile,
			HasTouch:          &emulateTemplate.HasTouch,
			UserAgent:         &emulateTemplate.UserAgent,
			Viewport:          emulateTemplate.Viewport,
		}

		context, err := browser.NewContext(newContextOptions)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		err = context.Route("**/*", routeHandler)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		page, err := context.NewPage()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		_, err = page.Goto(urlParam, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		for _, script := range scripts {
			_, err := page.AddScriptTag(playwright.PageAddScriptTagOptions{Content: &script})
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		}

		_, err = page.Evaluate(signatureFunctionsScript)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		verifyFP := "verify_5b161567bda98b6a50c0414d99909d4b"

		newURL := urlParam + "&verifyFp=" + verifyFP

		signature, err := page.Evaluate(fmt.Sprintf(`generateSignature("%s")`, newURL))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		signedURL := newURL + "&_signature=" + signature.(string)

		u, err := url.Parse(signedURL)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		queryString := u.RawQuery

		generatedBogus, err := page.Evaluate(fmt.Sprintf(`generateBogus("%s", "%s")`, queryString, emulateTemplate.UserAgent))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		signedURL += "&X-Bogus=" + generatedBogus.(string)

		encrypt_password := "webapp1.0+202106"
		queryString += "&is_encryption=1"

		password := padRight(encrypt_password, "\x00", 16)
		xxttparams, _ := xttparams(queryString, password)
		xxttparams += ""

		response := fiber.Map{
			"signature":  signature.(string),
			"signed_url": signedURL,
			"xxttparams": xxttparams,
			"verify_fp":  verifyFP,
			"user_agent": emulateTemplate.UserAgent,
		}

		return c.JSON(response)
	})

	app.Listen(":3000")
}

func padRight(str, pad string, length int) string {
	for len(str) < length {
		str += pad
	}
	return str
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func xttparams(queryStr, password string) (string, error) {
	queryStr += "&is_encryption=1"
	paddedPassword := password + string(make([]byte, 16-len(password)))
	queryBytes := []byte(queryStr)
	paddedQueryBytes := pad(queryBytes, aes.BlockSize)

	block, err := aes.NewCipher([]byte(paddedPassword))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(paddedQueryBytes))
	iv := ciphertext[:aes.BlockSize]
	copy(iv, []byte(paddedPassword)[:aes.BlockSize])

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedQueryBytes)

	return base64.StdEncoding.EncodeToString(ciphertext[aes.BlockSize:]), nil
}
