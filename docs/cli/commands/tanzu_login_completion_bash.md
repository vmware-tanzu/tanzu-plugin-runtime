## tanzu login completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(login completion bash)

To load completions for every new session, execute once:

#### Linux:

	login completion bash > /etc/bash_completion.d/login

#### macOS:

	login completion bash > $(brew --prefix)/etc/bash_completion.d/login

You will need to start a new shell for this setup to take effect.


```
tanzu login completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [tanzu login completion](tanzu_login_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 14-Sep-2022