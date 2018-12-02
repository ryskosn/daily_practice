# learn OCaml

## install

install OCaml

```sh
$ sudo port -N install opam
$ opam init
$ opam switch list -all
$ opam switch 4.06.0
```

### dune

install dune

```sh
$ opam update
$ opam install dune
```

- [ocaml/dune: A composable build system for OCaml](https://github.com/ocaml/dune)
- [dune documentation](https://jbuilder.readthedocs.io/en/latest/index.html#)
- [Jane Street Open Source](https://opensource.janestreet.com/)


dune file
```lisp
(executable
    (name foo))
```

```sh
$ dune build foo.exe
```

```sh
$ tree
.
├── README.md
├── _build
│   ├── default
│   │   ├── foo.exe
│   │   └── foo.ml
│   └── log
├── dune
├── dune-project
└── foo.ml
```

```sh
$ dune exec ./foo.exe
```

## tips

### tuareg-mode

- `C-c C-t` : show type of expression at cursol

- [ocaml/tuareg: Emacs OCaml mode](https://github.com/ocaml/tuareg)

## ref

- [「OCAML 入門」入門 - Qiita](https://qiita.com/kaizen_nagoya/items/456bedf9f68b512663da)
