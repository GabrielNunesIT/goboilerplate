# Go Boilerplate

Go Boilerplate is a CLI tool to bootstrap your Go projects with a predefined set of folders and files. It prompts the user for some basic information and creates a project structure based on user input.

## Features

- **Interactive CLI**: Prompts the user for project details.
- **Customizable Templates**: Use embedded templates to define your project's initial structure.
- **Seamless Integration**: Works with `go install` to make it easy to distribute and install.

## Installation

Install Go Project Builder using the `go install` command:

```sh
go install github.com/GabrielNunesIT/goboilerplate@latest
```

## Usage
```sh
goboilerplate
```

## Example
```sh
goboilerplate

What do you want to name the app?
Example

Where do you want to create the project? (Default is in the current folder.)
Press ENTER for default.
./example_project

Will your project include an API? [y/N]
n

Will your project include assets like images, logos, etc... ? [y/N]
n

Will your project contain Githooks? [y/N]
n

##################################################
Project creation finalized!!
##################################################
```

The layout produced is based on https://github.com/golang-standards/project-layout.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
