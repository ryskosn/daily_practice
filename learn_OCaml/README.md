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


#### dune file

```lisp
(executable
    (name foo))
```

```lisp
(library
 (name learn_ocaml)
 ;; (inline_tests (flags (-verbose)))
 (inline_tests)
 (preprocess (pps ppx_inline_test)))
```

`executable` か `library` を指定する。

`ppx_inline_test` は `library` に対してのみ有効。

- [error: unknown field inline_tests (conclusion: not a bug) · Issue #745 · ocaml/dune](https://github.com/ocaml/dune/issues/745)

#### build & execute

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

#### inline test

```ocaml
(* purpose: 原点から受け取った座標 pair までの距離を求める *)
(* kyori : float * float -> float *)
(* pattern: match pair with (a, b) *)

let kyori pair = match pair with
    (a, b) -> sqrt (a *. a +. b *. b)

(* test *)
let%test "test4" = kyori (3.0, 4.0) = 5.0
let%test "test5" = kyori (5.0, 12.0) = 13.0
let%test "test6" = kyori (8.0, 15.0) = 17.0
;;
```

以下コマンドで実行する。

```sh
$ dune runtest
```

`dune` ファイルで `flag` を指定できる。`-verbose` にするとテスト成功時にも詳細を表示する。
最初は `-verbose` にしていたが少々うるさいので変更した。

```lisp
(inline_tests (flags (-verbose)))
```

#### watch mode

`dune build`、`dune runtest` はオプション `-w` をつけるとファイルの変更を検知して自動的に再度ビルド or テストしてくれる。

ファイルを監視する `inotifywait` か `fswatch` が必要。macOS の場合、`fswatch` でよいらしい。

```sh
$ sudo port install fswatch
```

- [Watch mode — dune documentation](https://dune.readthedocs.io/en/latest/usage.html?highlight=watch#watch-mode)


## tips

### tuareg-mode

- `C-c C-t` : show type of expression at cursol

- [ocaml/tuareg: Emacs OCaml mode](https://github.com/ocaml/tuareg)

## ref

- [「プログラミングの基礎」を使った授業紹介](http://pllab.is.ocha.ac.jp/~asai/book-mov/)
- [「OCAML 入門」入門 - Qiita](https://qiita.com/kaizen_nagoya/items/456bedf9f68b512663da)
