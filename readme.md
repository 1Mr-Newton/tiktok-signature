
---

# TikTok Signature Generator API

This project provides an API to generate a TikTok signature and additional required parameters for requests. It leverages Playwright for browser automation and Fiber for the web framework.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoint](#api-endpoint)
- [Credits](#credits)
- [License](#license)

## Overview

This project creates an API that generates the TikTok signature and other necessary parameters to authenticate requests. It uses Playwright to automate the process of visiting a TikTok URL and executing JavaScript to generate the signature.

## Features

- Generates TikTok signature for a given URL
- Uses Playwright for headless browser automation
- Flexible configuration and easy-to-use API
- Built with Fiber, a fast HTTP web framework in Go

## Requirements

- Go (latest version recommended)
- Playwright for Go
- Fiber web framework

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/1mr-newton/tiktok-signature.git
cd tiktok-signature
```

2. **Install dependencies:**

```bash
go get -u github.com/gofiber/fiber/v2
go get -u github.com/playwright-community/playwright-go
```

3. **Download and install Playwright:**

```bash
playwright install
```

## Usage

1. **Prepare the scripts:**
   Ensure you have the necessary scripts (`bogus.txt`, `navigator.txt`, `signature_functions.txt`, `signer.txt`, `webmssdk.txt`) in the `cmd/scripts` directory.

2. **Run the API server:**

```bash
go run main.go
```

3. **Make a request to the API:**
   Use a tool like `curl` or Postman to make a GET request to the API.

`Example URL`:  ```https://www.tiktok.com/api/post/item_list/?WebIdLastTime=1714421730&aid=1988&app_language=en&app_name=tiktok_web&browser_language=en-US&browser_name=Mozilla&browser_online=true&browser_platform=MacIntel&browser_version=5.0%20%28Macintosh%3B%20Intel%20Mac%20OS%20X%2010_15_7%29%20AppleWebKit%2F537.36%20%28KHTML%2C%20like%20Gecko%29%20Chrome%2F126.0.0.0%20Safari%2F537.36&channel=tiktok_web&cookie_enabled=true&count=35&coverFormat=2&cursor=0&device_id=7363385157666391557&device_platform=web_pc&focus_state=true&from_page=user&history_len=5&is_fullscreen=false&is_page_visible=true&language=en&odinId=7363384667837826054&os=mac&post_item_list_request_type=0&priority_region=&referer=https%3A%2F%2Fwww.tiktok.com%2Fforyou&region=GH&root_referer=https%3A%2F%2Fwww.tiktok.com%2Fforyou&screen_height=1080&screen_width=1920&secUid=MS4wLjABAAAA1UWBntm1n1BFYlyVP4D7ddbfy_7I2sFo9X67s347pignH3dPyJcn7XFzEsmi4l0Z&tz_name=Africa%2FAccra&userId=6805706310416925702&verifyFp=verify_lvlehb8u_stK0hmVj_vwcR_4NVX_9qel_yBllzpPFajdT&webcast_language=en&msToken=LFGH8L4L05k6FN5aqczU1_tyF3JUNP5y5Pi9r3RSbMe_VlMJ_u7Cbpi88Vt9mLXZcLv5otsz5VvJubZBjrF3MXdGkahYitleoS-l3EadylYuXrWxjUHX5wPPMUTNE33PsYp9S76aJeJbzg==&X-Bogus=DFSzswVErG2ANaLjtWAPkHBeKL5E&_signature=_02B4Z6wo00001PqGkMAAAIDCRwm7KzaNr.D6hpRAAFjYa9```

```bash
curl "http://localhost:3000/sign?url=<your url here>"
```

## API Endpoint

### GET /sign

Generates a TikTok signature and other parameters for the given URL.

**Query Parameters:**

- `url` (string): The TikTok URL for which to generate the signature.

**Response:**

- `signature` (string): The generated TikTok signature.
- `signed_url` (string): The signed URL with the signature.
- `xxttparams` (string): The encrypted parameters.
- `verify_fp` (string): The verification fingerprint.
- `user_agent` (string): The user agent used for the request.

**Example:**

```json
{
  "signature": "_02B4Z6wo00f01tsCrUAAAIBBmh3r9C7VJ6bbEqnAANDA07",
  "signed_url": "https://www.tiktok.com/api/post/item_list/?WebIdLastTime=1714421730&aid=1988&app_language=en&app_name=tiktok_web&browser_language=en-US&browser_name=Mozilla&browser_online=true&browser_platform=MacIntel&browser_version=5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36&channel=tiktok_web&cookie_enabled=true&count=35&coverFormat=2&cursor=0&device_id=7363385157666391557&device_platform=web_pc&focus_state=true&from_page=user&history_len=5&is_fullscreen=false&is_page_visible=true&language=en&odinId=7363384667837826054&os=mac&post_item_list_request_type=0&priority_region=&referer=https://www.tiktok.com/foryou&region=GH&root_referer=https://www.tiktok.com/foryou&screen_height=1080&screen_width=1920&secUid=MS4wLjABAAAA1UWBntm1n1BFYlyVP4D7ddbfy_7I2sFo9X67s347pignH3dPyJcn7XFzEsmi4l0Z&tz_name=Africa/Accra&userId=6805706310416925702&verifyFp=verify_lvlehb8u_stK0hmVj_vwcR_4NVX_9qel_yBllzpPFajdT&webcast_language=en&msToken=LFGH8L4L05k6FN5aqczU1_tyF3JUNP5y5Pi9r3RSbMe_VlMJ_u7Cbpi88Vt9mLXZcLv5otsz5VvJubZBjrF3MXdGkahYitleoS-l3EadylYuXrWxjUHX5wPPMUTNE33PsYp9S76aJeJbzg==&X-Bogus=DFSzswVErG2ANaLjtWAPkHBeKL5E&_signature=_02B4Z6wo00001PqGkMAAAIDCRwm7KzaNr.D6hpRAAFjYa9&verifyFp=verify_5b161567bda98b6a50c0414d99909d4b&_signature=_02B4Z6wo00f01tsCrUAAAIBBmh3r9C7VJ6bbEqnAANDA07&X-Bogus=DFSzswSLqRsANHkjtWAPet9WcBJm",
  "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
  "verify_fp": "verify_5b161567bda98b6a50c0414d99909d4b",
  "xxttparams": "KgMc0joYXsLFgytpCAonUnX387UwJQR+YPrv3tJVo6Dn+rVRIGxnvVUTeIzYiAIb0I2Gv8tWdvPRAcM/dv9OvYBjy89gdKSghTmlodIBZ5qRiTc2hwLHSbTnNYnSOaoMHpu3v49Y8mB0kAFZMDwiHZkvtxlLSu1DKrI+ROvTsgisrk8uPA0OnmI7VR0eFQR0Fv4xQIUIxtc17dkxFqeyFJnSwmOsOONr5ntflYuEoD5zmSz7dp/U6Ib5wAFYLTL71XAA3gi773RnS5w4JVAr8WGxTjTq3rYJiqdInvgUvVcvo4m93yreuJ+XLTyJZLfpdAMSh6nV1czCiFMPrUIwYNSHWmMA0nQSOKSpqfvg/0M81F+Uw7r4nRLd8smEJKXU7Cg0aXCfEhWFowRUdmKZuY5dUdU/FzR5QGem0SC+8EXMOuTejKmw8sEDXwe1+h9/400U2q0ucX+86lMHy5MEiVFuZv8D0+dTJ9SdOt5+okVgOLSzpNgIinGVcysXP8VuwpeiPlCv03iOZGaG1YQMXSCgRPMJmbQTe9rsNgxBBB/hTlHXiUjs6j08YSZpm7O6kAN31H5Thg9DH9ZEaH27mBuDRhY+FgLklqjSZDIodUWZlnj/5xsq1ZbKEObvyir+z9xzqNpRcoZyS7VAfEwcVdbAdlg9yoJ49JCLDFJK4ewSVmrSQVrQaHIVw/eeOlieRTcQuC1pFs/yzeUb6Bh580QKLu1TZAUGGhuc4v+IuqnyYkJ+QPdda17HwBsdh8UPFjybOxVVOSk+X+3IwWubONM3G+FH/zqAvifsiuIpaX5amC5dCVhwC6IM2e/CJuvoNos/QfzE4mvalAY7CGPehgWHLJPmUCQnhiYkQ2imihJ+cuPVprbI7hMHdLbn6yI4w0TU0NmGzOdj9pkJJcJaQjKmXUE05Cnrx6seekevEJ2dQqMAXZgpRjB5eb8HasFo/TRFkGHDIw8EMvX1/2bl8ExxTa77rPZyuOz8pjDiGQhQbZACM2ShFNCqufHdCJ9j47vVUBDYA2LmYuVXbkrxXGfvQ1RwYU5effPTTi5xmFxFBkRkYOEOAn5AVHR1C0k1Wiub9Wps9C/AHDJx8c/frKEP4CQwU2z751uL5aQq2mfw93JtFxdabmg2YGezdSt4VpdkO/8KjpvwO3sKVmAJsA8APhFGZqqL4+V2O0CtexfNL12TT01RJSolcPrQ7Ae7krWrhy86Dj3ntCQ8gz90UAmLJeoQroK3Yxb3gUt0QQYxdYKJXDBWUIr4/FZ7BTlZhsMIGRXTwQtMyvz1QH+k9zyU5W2WRYWcyG2JVDHRx8SpOkio8G2KO/law6gpg4rh6zWr1FAdAZ3D+Si/Q96yy8vyMPZ+uLzLM6NSvh6/o4ft5C4kYPvLApgbm1/XhG9FAUQmADhDQKDxFS/AJq6y8FDwcEoErB85Osa5RE6ULTOzQTTO1M6c+iWGfO4f54p0eAr/M0MfowFJJy2na2Hx7todESrwTDJ34Jof0McPZDz11yeUvIAkUGSk9578XoONo0ad5WXB7sDs0LZmH/pFMkkynBtZs2VA2UUVSgBNM8m8USBs8R0fuxwAT9RLevthdB7c3vFDaJjHlLM1F+mgmC7/QouyrIgfIXZ2qDAoxzi9ELtB+rWATX6FH54Gja9jQvKCVzk96V5joeBE3lHX7Cku+6hwUgEVURuw76BeOugziVv39Le47gFx3vZafOpi7A8xGyDb4JCHoJiEvyWn00rE3sJdtf/DCX2zNeSgKRU="
}
```

## Credits

This project is inspired by [TikTok Signature Generator](https://github.com/carcabot/tiktok-signature) by [carcabot](https://github.com/carcabot). The idea and initial implementation were adapted from the JavaScript version to Go.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to modify this README as necessary to better fit your project's specifics.