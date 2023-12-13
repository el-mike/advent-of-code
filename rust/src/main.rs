mod common;
mod year_2021;
mod year_2023;

use std::env;
use std::time::Instant;
use crate::year_2023::day_04::scratchcards;

fn main() {
    let test_run = is_test_run();
    let start = Instant::now();

    scratchcards::run(test_run).unwrap();

    let duration = start.elapsed();

    println!("Execution time: {:?}", duration);
}

fn is_test_run() -> bool {
    let args: Vec<_> = env::args().collect();

    if args.len() < 2 { false } else { args[1] == "test" }
}
