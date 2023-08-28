use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("input.txt").expect("Fail to open file");
    let reader = BufReader::new(file);

    let mut total_score1: u32 = 0;
    let mut total_score2: u32 = 0;

    for line in reader.lines() {
        let l = line.expect("Fail to read line");

        let line_split: Vec<String> = l.split_whitespace().map(|s| s.to_string()).collect();

        let a = action_to_score1(line_split.get(0).expect("fail to get value of a"));
        let b = action_to_score1(line_split.get(1).expect("fail to get value of b"));

        total_score1 += cal_score(a, b);

        let b_action = action_to_score2(a, b);

        total_score2 += cal_score(a, b_action);
    }

    println!("{}", total_score1);
    println!("{}", total_score2);
}

fn action_to_score2(a: u32, b: u32) -> u32 {
    if b == 2 {
        return a;
    }

    if b == 1 {
        match a {
            1 => return 3,
            2 => return 1,
            _ => return 2,
        }
    }

    if a < 3 {
        return a + 1;
    }

    1
}

fn action_to_score1(action: &str) -> u32 {
    match action {
        "A" | "X" => 1,
        "B" | "Y" => 2,
        _ => 3, // C and Z
    }
}

fn cal_score(a: u32, b: u32) -> u32 {
    if a == b {
        return b + 3;
    }

    if (a == 1 && b == 3) || (a == 2 && b == 1) || (a == 3 && b == 2) {
        return b;
    }

    b + 6
}
