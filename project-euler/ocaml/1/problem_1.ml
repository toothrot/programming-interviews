let rec range a b =
  if a > b then []
  else a :: range (a + 1) b;;

let ismod35 x =
  x mod 3 = 0 || x mod 5 = 0;;

let sum35 (l : int list) : int =
  List.fold_left (fun acc x -> if ismod35 x then acc + x else acc ) 0 l;;

let foo = 
  range 0 999 in
    let answer = 
      sum35 foo in
        print_int answer;;
