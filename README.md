# makeproject

**WIP**: Command line application that easily creates a project that supports sensible settings for starter code. Features templates!

## Features

- [Gitpod](https://gitpod.io)
- [TravisCI](https://travis-ci.org)

## How to Use

```sh
NAME:
   makeproject - Create new projects easily.

USAGE:
   makeproject [global options] command [command options] [arguments...]

VERSION:
   0.0.0

AUTHOR:
   Dani Roxberry <dani@bitoriented.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value   --name=BEW-1.2-Project-Starter
   --type value   --type=python
   --help, -h     show help
   --version, -v  print the version
```

<!--
### Encrypting Environment Variables

Encrypt environment variables with the public key attached to your repository using the `travis` gem:

1. If you do not have the travis gem installed, run `gem install travis`.
2. In your repository directory: `travis encrypt MY_SECRET_ENV=super_secret --add env.global`
-->

## Contributing

Please fork the project, then submit a pull request to contribute!

### Add a New Project Template

Templates represent starter code that can be created for both languages (example: `python`) and frameworks (example: `flask`, `django`).

1. Add template files in the `templates` directory. Note the following:
   - Files must end with `.tmpl` to be included.
   - The `.tmpl` extension is stripped when the command is run.
2. Edit the `options/options.enum` file and add an option.
3. In the terminal, run: `enum -fp=options.enum`
4. Test `makeproject` with the new project: `makeproject --name=test_new --type=new_type`
