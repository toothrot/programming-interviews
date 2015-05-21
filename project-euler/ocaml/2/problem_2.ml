let rec fib x =
  match x with
  | 0 -> 1
  | 1 -> 1
  | x -> fib (x - 1) + fib (x - 2);;

let rec fibs a : int list =
  match a with
  | 0 -> [1]
  | 1 -> [1]
  | a -> fib a :: fibs (a - 1);;

let rec fibm curr max : int list =
  let head = List.hd (fibs curr) in
  if head > max then
    fibs (curr - 1)
  else
    fibm (curr + 1) max;;
  
let problem2 = 
  List.fold_left (fun acc x -> if x mod 2 == 0 then acc + x else acc) 0 (fibm 0 4000000);;

print_int problem2;;
