use std::mem;

fn main() {

    // 메모리 재할당
    let mut s = String::new();

    println!("{}", s.capacity());

    for _ in 0..5 {
        s.push_str("hello");
        println!("{}", s.capacity());
    }

    // 메모리 초기 할당
    let mut s = String::with_capacity(25);

    println!("{}", s.capacity());

    for _ in 0..5 {
        s.push_str("hello");
        println!("{}", s.capacity());
    }

    // ASCII 구조의 문자열
    let s1 = "hello"; // [104, 101, 108, 108, 111]
    assert_eq!(s1.len(), 5);

    let s2 = ['h', 'e', 'l', 'l', 'o'];
    let size: usize = s2.into_iter().map(|c| mem::size_of_val(&c)).sum(); // 배열 요소를 iterate하며 메모리 합함
    assert_eq!(size, 20);

    // ASCII가 아닌 Unicode 구조의 문자열
    let s1 = "안녕"; // [236, 149, 136, 235, 133, 149]
    assert_eq!(s1.len(), 6);

    let s2 = ['안', '녕'];
    let size: usize = s2.into_iter().map(|c| mem::size_of_val(&c)).sum();
    assert_eq!(size, 8);

    // 문자열 인덱싱
    let s1 = "안녕";
    let utf_str1 = s1.chars().nth(0).unwrap();
    let utf_str2 = s1.chars().nth(1).unwrap();

    println!("{}", utf_str1); // 안
    println!("{}", utf_str2); // 녕

    // 문자열 인덱싱 (byte 코드)
    let s1 = "안녕";
    let utf_str1 = std::str::from_utf8(&s1.as_bytes()[0..3]).unwrap();
    let utf_str2 = std::str::from_utf8(&s1.as_bytes()[3..6]).unwrap();

    println!("{}", utf_str1); // 안
    println!("{}", utf_str2); // 녕

    // &mut String
    let mut mutable_string = String::from("안녕") ;
    do_some_mutation(&mut mutable_string);
    println!("{}", mutable_string); // 안녕하세요
}

fn do_some_mutation(input: &mut String) {
    input.push_str("하세요"); 
}
