blsd
====

List directories in breadth-first order.

Usage
-----

```sh
blsd [DIRS...]
```

Installation
------------

```sh
go get -u github.com/junegunn/blsd
```

Using it with fzf
-----------------

```sh
command -v blsd > /dev/null && export FZF_ALT_C_COMMAND='blsd'
```

License
-------

MIT
