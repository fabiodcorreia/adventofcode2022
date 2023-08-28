use std::{
    fs::File,
    io::{BufRead, BufReader},
};

fn main() {
    // let file = File::open("example.txt").expect("Fail to open file");
    let file = File::open("input.txt").expect("Fail to open file");
    let reader = BufReader::new(file);

    let mut counter: u32 = 0;
    let mut result: Vec<u32> = Vec::new();

    for line in reader.lines() {
        let l = line.expect("Fail to read line");

        if l.is_empty() {
            result.push(counter);
            counter = 0;
            continue;
        }
        counter += l.parse::<u32>().expect("Fail to parse number")
    }

    result.push(counter);

    if let Some(max) = result.iter().max() {
        println!("Max = {}", max);
    }

    result.sort();

    let top3_cal: u32 = result.iter().rev().take(3).sum();

    println!("Top3= {}", top3_cal);
}
