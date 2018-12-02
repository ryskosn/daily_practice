(* Q 4.6 *)
(* purpose: 鶴の数から脚の数を求める *)
(* tsuru_no_ashi : int -> int *)
(* pattern: nothing *)

let tsuru_no_ashi n = n * 2

(* test *)
let%test "test" = tsuru_no_ashi 3 = 6
let%test "test" = tsuru_no_ashi 0 = 0
let%test "test" = tsuru_no_ashi 12 = 24
;;

(* Q 4.6.2 *)
(* purpose: 亀の数から脚の数を求める *)
(* kame_no_ashi : int -> int *)
(* pattern: nothing *)

let kame_no_ashi n = n * 4

(* test *)
let%test "test" = kame_no_ashi 2 = 8
let%test "test" = kame_no_ashi 0 = 0
let%test "test" = kame_no_ashi 16 = 64
;;

(* Q 4.7 *)
(* purpose: 鶴の数と亀の数から脚の数の合計を求める *)
(* tsurukame_no_ashi : int -> int -> int *)
(* pattern: nothing *)

let tsurukame_no_ashi a b =
  (tsuru_no_ashi a) + (kame_no_ashi b)

(* test *)
let%test "test" = tsurukame_no_ashi 2 4 = 20
let%test "test" = tsurukame_no_ashi 0 3 = 12
let%test "test" = tsurukame_no_ashi 5 0 = 10
;;
