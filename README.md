# Tanuki-Ledger

:construction: **WIP** :construction:

This application is right now in an early stage of development.

More descriptions coming soon.

## Development

This code follows the guidelines from (Fyne Conf 2022)[https://conf.fyne.io/archive/2022/BestPractice.pdf].

### Git-Hooks

All git-hooks are stored in the `git-hooks` directory. Please make sure they are executable and symlink them into `.git/hooks`.

Please note, that when creating the symlink, you have (adjust the origin path)[https://stackoverflow.com/a/4594681/18709964] in order for git to execute the script.

For the `pre-commit` hook you'll need to install `golangci-lint`, `goimports` and `gocyclo`. All are available trough the `go install` command. If you are using the devcontainer they are installed by default.