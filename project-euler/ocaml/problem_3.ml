let rec range a b =
  if a > b then []
  else a :: range (a + 1) b;;

let rec isPrime x =
  List.exists (fun i -> x mod i = 0) (range 0 (x/2));;
