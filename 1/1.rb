puts (0...1000).inject(0){|sum,n| (n % 5 == 0 || n % 3 ==0) ? sum+n : sum}

