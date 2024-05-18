# browser [![MIT](https://img.shields.io/github/license/dineshgowda24/browser)](https://github.com/dineshgowda24/browser/blob/main/LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/dineshgowda24/browser.svg)](https://pkg.go.dev/github.com/dineshgowda24/browser) [![Go report card](https://goreportcard.com/badge/github.com/dineshgowda24/browser)](https://goreportcard.com/report/github.com/dineshgowda24/browser) [![Build Status](https://dl.circleci.com/status-badge/img/circleci/MQTLZJuBejHgr2yqrojz3u/5NTLeuQeViQw2JaPQf7gKa/tree/main.svg?style=shield&circle-token=ab7a417fe410b8387c767f83568f7d2f2788ac4f)](https://dl.circleci.com/status-badge/redirect/circleci/MQTLZJuBejHgr2yqrojz3u/5NTLeuQeViQw2JaPQf7gKa/tree/main) [![Coverage](https://codecov.io/gh/dineshgowda24/browser/graph/badge.svg?token=XUA2VJW5FU)](https://codecov.io/gh/dineshgowda24/browser) [![X](https://img.shields.io/twitter/follow/_dineshgowda)](https://twitter.com/_dineshgowda)

<p align="center">
  <img src="logo.png">
</p>

## Why?

> *Why not?*

You need to detect the browser, device and platform from the user agent string in many use cases. I have personally come across the following use cases:

1. You want to render different HTML for different browsers, devices or platforms.
2. You want to render OG tags for scraping bots and social media sites.
3. You want to log the browser, device or platform for your analytics.
4. You want your backend to behave differently when a Google bot crawls your site for SEO.

I wanted a relatively extensible package that I could use in all the above use cases. So, I decided to write this package.

## Inspiration

The ruby gem [fnando/browser](https://github.com/fnando/browser) inspires this package. I have used the gem in some of my previous projects and liked it. All the credit goes to the author of the ruby gem, who has done a great job.

## Documentation

For detailed documentation visit [browser.dineshgowda.com](https://browser.dineshgowda.com). It has adoption guides, usage, contributing guidelines and also the list of all the matchers and browsers supported.

## Usage

```go
go get github.com/dineshgowda24/browser
```

### Browser Detection

```go
b, err := browser.NewBrowser("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.210 Mobile Safari/537.36")
 if err != nil {
  // handle error
 }

 // browser level information
 fmt.Println(b.Name())           // Chrome
 fmt.Println(b.Version())        // 90.0.4430.210
 fmt.Println(b.ShortVersion())   // 90
 fmt.Println(b.IsBrowserKnown()) // true
 fmt.Println(b.IsChrome())       // true

 // device level information
 fmt.Println(b.Device().Name())      // Samsung SM-A205U
 fmt.Println(b.Device().IsTablet())  // false
 fmt.Println(b.Device().IsSamsung()) // true

 // platform level information
 fmt.Println(b.Platform().Name())         // Android
 fmt.Println(b.Platform().Version())      // 10
 fmt.Println(b.Platform().IsAndroidApp()) // false

 // bot level information
 fmt.Println(b.Bot().Name())  // ""
 fmt.Println(b.Bot().IsBot()) // false
```

### Bot Detection

```go
b, err := browser.NewBrowser("APIs-Google (https://developers.google.com/webmasters/APIs-Google.html)")
if err != nil {
   // handle error
}

// browser level information
fmt.Println(b.Name())           // Unknown Browser
fmt.Println(b.Version())        // 0.0
fmt.Println(b.ShortVersion())   // 0
fmt.Println(b.IsBrowserKnown()) // false
fmt.Println(b.IsUnknown())      // true

// bot level information
fmt.Println(b.Bot().Name())  // "APIs-Google"
fmt.Println(b.Bot().IsBot()) // true
fmt.Println(b.Bot().Why())   // *bots.Known
```

## Contributing

If you want to contribute to this project, please read the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## Issues

If you find any issues with this package, please raise an issue. I will fix it as soon as possible. Please read the `Contributing` section if you want to resolve the issue and contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
