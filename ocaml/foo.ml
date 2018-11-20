let hoge = 10
let rec fact n = if n = 1 then 1 else n * fact (n - 1)
let%test _ = fact 4 = 24
(* let () = print_int hoge; *)
