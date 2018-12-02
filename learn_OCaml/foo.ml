(* purpose: 所持金が与えられたとき 126 円のチョコレートをいくつ買えるかを求める *)
(* chocolate : int -> int *)

let chocolate n = if n < 126 then 0
                  else n / 126

(* test *)
let%test "test1" = chocolate 100 = 0
let%test "test2" = chocolate 252 = 2
let%test "test3" = chocolate 500 = 3
;;

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
