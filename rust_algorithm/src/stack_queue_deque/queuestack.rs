use std::collections::VecDeque;



pub fn solve(n:usize, is_queue:Vec<u32> , store:Vec<u32>, inputs:Vec<u32>) -> Vec<u32>{ // u32::MAX = 4_294_967_295u32
  
  let mut rtn:Vec<u32> = Vec::with_capacity(inputs.len()); // 이건 미리 메모리만 확보하는 것이지 길이가 정해지는 것이 아님
  
  let mut deque: VecDeque<u32> = VecDeque::new();

  for i in 0..n{
    if is_queue[i].eq(&0) { // eq는 참조끼리 비교
      deque.push_front(store[i]);
    }
  }
  for i in 0..inputs.len() {
    deque.push_back(inputs[i]);
    rtn.push(deque.pop_front().unwrap());
  }
  
  rtn
}


#[cfg(test)]
mod test {
    use super::*;
    
    #[test]
    fn test1() {
      let n = 4_usize;
      let is_queue:Vec<u32> = vec![0,1,1,0];
      let store:Vec<u32> = vec![1,2,3,4];
      let input: Vec<u32> = vec![2,4,7];
      let result = solve(n,is_queue, store, input);
      assert_eq!( result, [4,1,2])
    }

    #[test]
    fn test2() {
      let n = 4_usize;
      let is_queue:Vec<u32> = vec![1,1,1,1,1];
      let store:Vec<u32> = vec![1,2,3,4,5];
      let input: Vec<u32> = vec![1,3,5];
      let result = solve(n,is_queue, store, input);
      assert_eq!( result, [1,3,5])
    }
}