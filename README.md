# Browser: A browser detection in Go (Golang)

## Why?

I wanted a package to detect the browser, device and platform from the user agent string. I found nothing similar in Go (Golang), so I wrote my package.

### Inspiration

This package is inspired by the ruby gem [fnando/browser](https://github.com/fnando/browser). I wanted to have something similar in Go (Golang) to detect the browser and the platform from the user agent string but I couldn't find anything similar. So I decided to write my own package inspired by the ruby gem.

## Design

I have tried to keep the architecture as similar as possible to the ruby gem. I have followed the language conventions and the package structure of Go.

The package is divided into three parts:

- **browser**: This package contains the code to detect the browser.
- **device**: This package contains the code to detect the device from the user agent string.
- **platform**: This package contains the code to detect the platform from the user agent string.

### Browser

A `Matcher` interface defines a matching behaviour for a user agent string.

```go
type Matchers interface {
    Match() bool
}
```

`BrowserMatchers` interface matches the user agent string with the browser. Any new browser can be added by implementing the `BrowserMatchers` interface.

```go
type BrowserMatcher interface {
    Matcher
    Name() string        // Name returns the name of the browser.
    FullVersion() string // FullVersion returns the full version of the browser.
}
```

### Device

`DeviceMatchers` interface matches the user agent string with the device. Any new device can be added by implementing the `DeviceMatchers` interface.

```go
type DeviceMatcher interface {
    Matcher
    Name() string // Name of the device
}
```

### Platform

`PlatformMatchers` interface matches the user agent string with the platform. Any new platform can be added by implementing the `PlatformMatchers` interface.

```go
type PlatformMatcher interface {
    Matcher
    Name() string    // Name returns the name of the platform.
    Version() string // Version returns the version of the platform.
}
```
