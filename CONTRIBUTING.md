# Contributing

## Requirements

- [Go](https://golang.org/dl/) 1.19 or higher

## Getting Started

1. Fork the repository and clone it locally
2. Run `go mod download` to download dependencies
3. Run `go test ./...` to run all tests

## Testing

The tests are located in the `matchers`, `devices`, `platforms` and `browser` packages. The tests are using the `assets/test` directory to load the user agents data.

Package follows BDD style test cases. The tests are using the [GoConvey](https://github.com/smartystreets/goconvey) library to write BDD style tests.

## Adding a new browser support

Every browser implements the `BrowserMatcher` interface.

1. Add the file to the `matchers` package.
2. Add the new browser the list of matchers in `browser.go`. Make sure the order of the matchers is correct. The order of the matchers is important because the first matcher that returns `true` will be used.
3. Add specific tests for the new browser in the `matchers` package. Make sure to test all the possible user agents for the browser. You can use [this](https://developers.whatismybrowser.com/useragents/explore/) website to find user agents for the browser.
4. Add the user agents data to [matchers.yml](assets/test/matchers.yml) file.
5. Add all necerssary functions to `Browser` struct in `browser.go` file.

## Adding a new device support

Every device implements the `DeviceMatcher` interface.

1. Add the file to the `devices` package.
2. Add the new device to list of `DeviceMatcher` in `device.go`. Make sure the order of the matchers is correct. The order of the matchers is important because the first matcher that returns `true` will be used.
3. Add specific tests for the new device in the `devices` package. Make sure to test all the possible user agents for the device. You can use [this](https://developers.whatismybrowser.com/useragents/explore/) website to find user agents for the device.
4. Add the user agents data to [devices.yml](assets/test/devices.yml) file.
5. Add all necerssary functions to `Device` struct in `device.go` file.

## Adding a new platform support

Every platform implements the `PlatformMatcher` interface.

1. Add the file to the `platforms` package.
2. Add the new platform to list of `PlatformMatcher` in `platform.go`. Make sure the order of the matchers is correct. The order of the matchers is important because the first matcher that returns `true` will be used.
3. Add specific tests for the new platform in the `platforms` package. Make sure to test all the possible user agents for the platform. You can use [this](https://developers.whatismybrowser.com/useragents/explore/) website to find user agents for the platform.
4. Add the user agents data to [platforms.yml](assets/test/platforms.yml) file.
5. Add all necerssary functions to `Platform` struct in `platform.go` file.

## Adding a new known bot support

1. Add the new bot to [known_bots.yml](assets/internal/known_bots.yml).
