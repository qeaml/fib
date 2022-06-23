# fib

fib, short for **f**inal **i**mage **b**oard, is a simple, straight-to-the-point
image board software written in Go. This is the final iteration of QMLIB.

## Tech stack

fib is written in [Go][go], and uses the [Fiber][fiber] framework. It has been
written with the assistance of [GitHub Copilot][copilot]. Once the app is mature
enough, [Fauna DB][fauna] will be used to store the data.

## Roadmap

The following features are finished:

* Temporary storage via JSON.
* Near-final routing.

And the following is planned to be added as soon as possible:

* JSON API.
* User registration and authentication.
* Uploading images.
* Image moderation.
* Album creation and management.

Last but not least, the following is planned to be added in the future:

* Image search.
* Image voting.
* Image commenting.
* Friends.

## Installation

1. Clone the repository.

```sh
git clone https://github.com/qeaml/fib
cd fib
```

2. Create the config file. Check the [Configuration](#Configuration) section for
    more information.

```sh
# For example, if you wish to use Vim to edit the file straight away
vim config.naml
# Or if you prefer to use Emacs, you can use the following command:
emacs config.naml
# Eventually, create a blank config.naml file and edit it with your favorite editor
touch config.naml
```

3. Install the dependencies.

```sh
go get -v ./...
```

4. Build the application.

```sh
go build
```

5. Run the application.

```sh
./fib # or just 'fib' on Windows
```

## Troubleshooting

In case the application does not start, double-check that the configuration file
is correct and present in the directory. Also ensure that the `data` directory
is present.

## Configuration

fib is configured using a file called `config.naml`. This file is a
[NAML](https://github.com/naml-conf/naml) file. Below is a reference of all the
configuration options.

Name | Type | Description
-----|------|------------
`port` | integer | The port to listen on.

[go]: https://golang.org/
[fiber]: https://github.com/gofiber/fiber
[copilot]: https://github.com/features/copilot/
[fauna]: https://faunadb.com/
