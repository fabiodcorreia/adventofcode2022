use std::{
    fs::File,
    io::{BufRead, BufReader},
};

fn main() {
    part1();
    part2();
}

fn part2() {
    let file = File::open("input.txt").expect("Fail to open file");
    let reader = BufReader::new(file);
    let mut counter: u32 = 0;

    let mut lines = reader.lines().map(|line| line.expect("fail to read line"));

    while let Some(line1) = lines.next() {
        let line2 = lines.next().expect("Expected another line");
        let line3 = lines.next().expect("Expected another line");

        let badge_item = find_trip(&line1, &line2, &line3);
        counter += char_to_priority(badge_item);
    }
    println!("{counter}")
}

fn find_trip(ruck1: &str, ruck2: &str, ruck3: &str) -> char {
    for ch in ruck1.chars() {
        if ruck2.contains(ch) && ruck3.contains(ch) {
            return ch;
        }
    }
    '\0'
}

fn part1() {
    let file = File::open("input.txt").expect("Fail to open file");
    let reader = BufReader::new(file);
    let mut counter: u32 = 0;

    for line in reader.lines() {
        let l = line.expect("Fail to read line");
        let middle = l.len() / 2;
        if let Some(dup_val) = find_dup(&l[..middle], &l[middle..]) {
            counter += char_to_priority(dup_val);
        }
    }
    println!("{counter}")
}

fn find_dup(comp1: &str, comp2: &str) -> Option<char> {
    comp1.chars().find(|&item| comp2.contains(item))
}

fn char_to_priority(c: char) -> u32 {
    match c {
        'a'..='z' => c as u32 - 'a' as u32 + 1,
        'A'..='Z' => c as u32 - 'A' as u32 + 27,
        _ => 0,
    }
}
