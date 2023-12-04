mod stack_queue_deque;
use std::io;

fn main() {

    let buf = io::read_to_string(io::stdin()).unwrap();
    let mut input  = buf.split_ascii_whitespace();
    
    let n:usize = input.next().unwrap().parse().unwrap();
    let is_queue:Vec<u32> = (0..n).map(|_| input.next().unwrap().parse().unwrap()).collect();
    let store:Vec<u32> = (0..n).map(|_| input.next().unwrap().parse().unwrap()).collect();
    let m:usize = input.next().unwrap().parse().unwrap();
    let input:Vec<u32> = (0..m).map(|_| input.next().unwrap().parse().unwrap()).collect();

    let rtn = stack_queue_deque::queuestack::solve(n, is_queue, store, input);
    for e in rtn {
        print!("{} ", e);
    }
}
