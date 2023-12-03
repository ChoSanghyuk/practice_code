mod stack_queue_deque;
use std::io;

fn main() {
    
    let mut n = String::new();
    io::stdin().read_line(&mut n).unwrap();
    let n:usize = n.trim().parse().unwrap();

    let mut is_queue = String::new();
    io::stdin().read_line(&mut is_queue).unwrap();
    let is_queue:Vec<u32> = is_queue.trim().split_whitespace().map(|s| s.parse().unwrap()).collect();

    let mut store= String::new();
    io::stdin().read_line(&mut store).unwrap();
    let store:Vec<u32> = store.trim().split_whitespace().map(|s| s.parse().unwrap()).collect();

    let mut input= String::new();
    io::stdin().read_line(&mut input).unwrap();

    let mut input= String::new();
    io::stdin().read_line(&mut input).unwrap();
    let input:Vec<u32> = input.trim().split_whitespace().map(|s| s.parse().unwrap()).collect();

    let rtn = stack_queue_deque::queuestack::solve(n, is_queue, store, input);
    for e in rtn {
        print!("{} ", e);
    }
}
