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
	defaultURL := "https://www.tiktok.com/@rihanna?lang=en"

	bogusScript, err := readFile("cmd/scripts/bogus.txt")
	if err != nil {
		panic(err)
	}
	navigatorScript, err := readFile("cmd/scripts/navigator.txt")
	if err != nil {
		panic(err)
	}
	signatureFunctionsScript, err := readFile("cmd/scripts/signature_functions.txt")
	if err != nil {
		panic(err)
	}
	signerScript, err := readFile("cmd/scripts/signer.txt")
	if err != nil {
		panic(err)
	}
	webmssdkScript, err := readFile("cmd/scripts/webmssdk.txt")
	if err != nil {
		panic(err)
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
		panic(err)
	}

	browser, err := pw.Chromium.Launch(options)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	err = context.Route("**/*", routeHandler)
	if err != nil {
		panic(err)
	}

	page, err := context.NewPage()
	if err != nil {
		panic(err)
	}

	_, err = page.Goto(defaultURL, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle})
	if err != nil {
		panic(err)
	}

	for _, script := range scripts {
		_, err := page.AddScriptTag(playwright.PageAddScriptTagOptions{Content: &script})
		if err != nil {
			panic(err)
		}
	}

	_, err = page.Evaluate(signatureFunctionsScript)
	if err != nil {
		panic(err)
	}

	verifyFP := "verify_5b161567bda98b6a50c0414d99909d4b"
	link := "https://www.tiktok.com/api/post/item_list/?WebIdLastTime=1714421730&aid=1988&app_language=en&app_name=tiktok_web&browser_language=en-US&browser_name=Mozilla&browser_online=true&browser_platform=MacIntel&browser_version=5.0%20%28Macintosh%3B%20Intel%20Mac%20OS%20X%2010_15_7%29%20AppleWebKit%2F537.36%20%28KHTML%2C%20like%20Gecko%29%20Chrome%2F126.0.0.0%20Safari%2F537.36&channel=tiktok_web&cookie_enabled=true&count=35&coverFormat=2&cursor=0&device_id=7363385157666391557&device_platform=web_pc&focus_state=true&from_page=user&history_len=3&is_fullscreen=false&is_page_visible=true&language=en&odinId=7363384667837826054&os=mac&region=GH&screen_height=1080&screen_width=1920&secUid=MS4wLjABAAAAbhSKTPb2GvF2GnRmj6Ai9AlzV2A6PV69lE2LlhjM3yGv153htSAD3PBoht2j8kdN&tz_name=Africa%2FAccra&verifyFp=verify_lvlehb8u_stK0hmVj_vwcR_4NVX_9qel_yBllzpPFajdT&webcast_language=en&msToken=VVUO_HrEYchrqKjhZewtXVB2ad51q3bdQD1LLhZ0_FfJhyoEKu1mh_rWgyDAscPDC6_kmwsB5kXXHPg1MOqSjD2186ncdI_1XSVJB14HWvR8ni7reph-6Z1PxZwVRyOkQKPR0BHr7MGJbA==&X-Bogus=DFSzswVOYHsANJoVtWlY6qBeKLRW&_signature=_02B4Z6wo000019PpFTwAAIDBbmY-1dYIjiPT6RGAAJKSd9&verifyFp=verify_5b161567bda98b6a50c0414d99909d4b&_signature=_02B4Z6wo00f01bAECdQAAIBC8RtPYbQIutWwFA1AAApzeb&X-Bogus=DFSzswSLptXANHkjtWAbIU9WcBjp&cursor=1713785978000"

	newURL := link + "&verifyFp=" + verifyFP

	signature, err := page.Evaluate(fmt.Sprintf(`generateSignature("%s")`, newURL))
	if err != nil {
		panic(err)
	}

	signedURL := newURL + "&_signature=" + signature.(string)

	u, err := url.Parse(signedURL)
	if err != nil {
		panic(err)
	}
	queryString := u.RawQuery

	generatedBogus, err := page.Evaluate(fmt.Sprintf(`generateBogus("%s", "%s")`, queryString, emulateTemplate.UserAgent))
	if err != nil {
		panic(err)
	}

	signedURL += "&X-Bogus=" + generatedBogus.(string)

	encrypt_password := "webapp1.0+202106"
	queryString += "&is_encryption=1"

	password := padRight(encrypt_password, "\x00", 16)
	xxttparams, _ := xttparams(queryString, password)
	xxttparams += ""

	fmt.Println("Signature: ", signature.(string)+"\n\n")
	fmt.Println("Signed URL: ", signedURL+"\n\n")
	fmt.Println("XXTParams: ", xxttparams+"\n\n")
	fmt.Println("verify_fp: ", verifyFP+"\n\n")
	fmt.Println("User Agent: ", emulateTemplate.UserAgent+"\n\n")

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
