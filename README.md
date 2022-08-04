# 🦄 arcadia: ___Go template for Cobra based cli applications___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/arcadia.svg)](https://pkg.go.dev/github.com/snivilised/arcadia)
[![Go report](https://goreportcard.com/badge/github.com/snivilised/arcadia)](https://goreportcard.com/report/github.com/snivilised/arcadia)
[![Coverage Status](https://coveralls.io/repos/github/snivilised/arcadia/badge.svg?branch=master)](https://coveralls.io/github/snivilised/arcadia?branch=master&kill_cache=1)
[![Arcadia Continuous Integration](https://github.com/snivilised/arcadia/actions/workflows/ci-workflow.yml/badge.svg)](https://github.com/snivilised/arcadia/actions/workflows/ci-workflow.yml)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" /></a>
</p>

## 🔰 Introduction

This project is a template to aid in the startup of Go cli applications.

## 🔨 Usage

## 🎀 Features

<p align="left">
  <a href="https://onsi.github.io/ginkgo/"><img src="https://onsi.github.io/ginkgo/images/ginkgo.png" width="100" /></a>
  <a href="https://onsi.github.io/gomega/"><img src="https://onsi.github.io/gomega/images/gomega.png" width="100" /></a>
</p>

+ unit testing with [Ginkgo](https://onsi.github.io/ginkgo/)/[Gomega](https://onsi.github.io/gomega/)
+ implemented with [🐍 Cobra](https://cobra.dev/) cli framework, assisted by [🐲 Cobrass](https://github.com/snivilised/cobrass)
+ i18n with [go-i18n](https://github.com/nicksnyder/go-i18n)
+ linting configuration and pre-commit hooks, (see: [linting-golang](https://freshman.tech/linting-golang/)).

## 🧰 Developer Info

By using this template, there is no need to use the cobra-cli to scaffold your application as this has been done already. It should be noted that the structure that is generated the cobra-cli has been significantly changed in this template, mainly to remove use of the __init()__ function and to minimise use of package level global variables. For a rationale, see [go-without-package-scoped-variables](https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables).

### 📝 Checklist of required changes

The following is list of actions that must be performed before using this template. Most of the changes concern changing the name `Arcadia` to the name of the new application.

As the template is instantiated from github, the new name will automatically replace the top level directory name, that being ___arcadia___.

+ `github actions workflow`: If the client application needs to use github actions for continuous integration, then the name of the [workflow](.github/workflows/ci-workflow.yml) needs to be changed. If not, then the workflow file should be deleted
+ `remove the dummy code`: __widget-cmd.go__, __greeting.go__ and its associated test __greeting_test.go__ (but only do this once new valid tests are ready to replace it, to avoid references being removed after _go mod tidy_)
+ `replace bootstrap testcase`: There is a test case which by default is set to invoke the __widget__ command. When the user is ready to remove this command, then the corresponding test case should be modified to invoke another command with appropriate parameters. This test case is there to ensure that the ___bootstrapping___ process works, as opposed to checking the validatity of the command itself.

+ `change ApplicationName`: modify to reflect the new paplication name. This application name is incorporated into the name of any translation files to be loaded.
+ `replace README content`
+ `review bootstrap.go`: this will need to be modified to invoke creation of any custom commands. The `execute` method of __bootstrap__ should be modified to invoke command builder. Refer to the `widget` command to see how this is done.
+ `update BINARY_NAME`: Inside _Taskfile.yml_, change the value of ___BINARY_NAME___ to the name of the client application.
+ `update translation file`: Inside _Taskfile.yml_, add support for loading any translations that the app will support. By default, it deploys a translation file for __en-US__ so this needs to be updated as appropriate.
+ `update email address in copyright statement`: The __root.go__ file contains a placeholder for an email address, update this comment accordingly.
+ `create .env file`: Add any appropriate secrets to a newly created .env in the root directory and to enable the __deploy__ task to work, define a __DEPLOY_TO__ entry that defines where builds should be deployed to for testing
+ `update message id`: This package supports i18n and as part of that defines messages that need to be translated. The user needs to update the message ids of defined messages in `messages.go`, which by default contain ___arcadia___ as part of the id.

### ㊗️ l10n Translations

This template has been setup to support localisation. The default language is `en-GB` with support for `en-US`. There is a translation file for `en-US` defined as __src/internal/l10n/out/arcadia.active.en-US.json__.

Make sure that the go-i18n package has been installed so that it can be invoked as cli, see [go-i18n](https://github.com/nicksnyder/go-i18n) for installation instructions.

To maintan localisation of the application, the user must take care to implement all steps to ensure translatablity of all user facing messages. Whenever there is a need to add/change user facing messages including error messages, to maintain this state, the user must:

- define template struct (__xxxTemplData__) in __src/internal/l10n/messages.go__ and corresponding __Message()__ method. All messages are defined here in the same location, simplifying the message extraction process as all extractable strings occur at the same place. Please see [go-i18n](https://github.com/nicksnyder/go-i18n) for all translation/pluralisation options and other regional sensitive content.
- define a corresponding helper function in __src/internal/translate/messages.go__. These helper functions are the ones that the rest of the application will use in order to generate region sensitive user facing string content.
- cd to the ___l10n___ at __src/internal/l10n/__
- run `goi18n extract -format json`, this will create an updated __active.en.json__ file
- run `goi18n merge -outdir out -format json active.en.json translate.en-US.json`
- rename __out/active.en-US.json__ to __out/arcadia.active.en-US.json__. The name __arcadia__ should be changed to the name of the new app (which should correspond with `ApplicationName` defined in __src/app/command/root-cmd.go__).

The file __out/arcadia.active.en-US.json__ is the translation file that will be deployed with the executable. Of course, if you want to use a config file format other than `json`, then there will be a little more work to do, but is fairly straight forward.

### 🧪 Quick Test

To check the app is working (as opposed to running the unit tests), build and deploy:

> task b

> task d

NB: the `deploy` task has been set up for windows by default, but can be changed at will.

Check that the executable and the US language file __arcadia.active.en-US.json__ have both been deployed. Then invoke the widget command with something like

> arcadia widget -p "P?\<date\>" -t 30

Optionally, the user can also specify the ___directory___ flag:

> arcadia widget -p "P?\<date\>" -t 30 -d foo-bar.txt

... where ___foo-bar.txt___ should be replaced with a file that actually exists.

Since the `widget` command uses `Cobrass` option validation to check that the file specified exists, the app will fail if the file does not exist. This serves as an example of how to implement option validation with `Cobrass`.
