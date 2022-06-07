blsd
====

List directories in breadth-first order.

Usage
-----

```sh
blsd [DIRS...]
```

Install
-------

```sh
bash <(curl -fL https://raw.githubusercontent.com/junegunn/blsd/master/install)
```

Or, you can use Homebrew to build blsd:

```sh
brew install --HEAD junegunn/tap/blsd
```

Build
-----

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
