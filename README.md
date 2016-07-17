blsd
====

List directories in breadth-first order.

Usage
-----

```sh
blsd [DIRS...]
```

Build
------------

```sh
make
```

Using it with fzf
-----------------

```sh
command -v blsd > /dev/null && export FZF_ALT_C_COMMAND='blsd'
```

License
-------

MIT
