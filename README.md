# Browser

![](logo.png)

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

## Design

I have kept the design as similar as possible to the gem but made changes where I felt necessary per the Go.

The following are the sub-packages:

- **matchers**: This package defines all the browser matchers.
- **devices**: This package defines all the device matchers.
- **platforms**: This package defines all the platform matchers.
- **bots**: This package defines all the bots matchers.

A `Matcher` interface defines a matching behaviour for a user agent string.

```go
type Matchers interface {
    Match() bool    // Match returns true if the user agent string matches the matcher.
    Name() string   // Name returns the name of the matcher.
}
```

### Matchers

`BrowserMatcher` interface matches the user agent string with the browser. Implement the `BrowserMatcher` interface to add a new browser.

```go
type BrowserMatcher interface {
    Matcher
    Version() string // Version returns the full version of the browser.
}
```

### Devices

`DeviceMatcher` interface matches the user agent string with the device. Implement the `DeviceMatcher` interface to add a new device.

```go
type DeviceMatcher interface {
    Matcher
}
```

### Platforms

`PlatformMatcher` interface matches the user agent string with the platform. Implement the `PlatformMatcher` interface to add a new device.

```go
type PlatformMatcher interface {
    Matcher
    Version() string // Version returns the version of the platform.
}
```

### Bots

`BotMatcher` interface matches the user agent string with the bot. Implement the `BotMatcher` interface to add a new bot.

```go
type BotMatcher interface {
    Matcher
}
```

### Browser Struct

`Browser` struct abstracts a lot of functionality. It uses the `BrowserMatcher`, `DeviceMatcher`, `PlatformMatcher` and `BotMatcher` interfaces to match the user agent string with the browser, device, platform and bot respectively. All the matchers are executed in the order they are defined in the `Browser` struct. The first matcher that returns `true` will be used.

A ton of helper functions are defined in the `Browser` struct to make it easy to use.

## Usage

```go
go get github.com/dineshgowda24/browser
```

```go
b, err := browser.New("Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.210 Mobile Safari/537.36")
if err != nil {
    // handle error
}

// browser level information
b.Name()            // Chrome
b.Version()         // 90.0.4430.210
b.ShortVersion()    // 90
b.IsBrowserKnown()  // true
b.IsChrome()        // true

// device level information
b.Device().Name()       // Samsung SM-A205U
b.Device().IsTablet()   // false
b.Device().IsSamsung()  // true

// platform level information
b.Platform().Name()         // Android
b.Platform().Version()      // 10
b.Platform().IsAndroidApp() // false

// bot level information(irrelevant for this user agent)
b.Bot().Name()      // ""
b.Bot().IsBot()     // false
```

## Contributing

If you want to contribute to this project, please read the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## Issues

If you find any issues with this package, please raise an issue. I will fix it as soon as possible. Please read the `Contributing` section if you want to resolve the issue and contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
