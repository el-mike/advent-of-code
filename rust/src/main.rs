mod common;
mod year_2021;

use std::env;
use std::time::Instant;
use crate::year_2021::day_04::giant_squid;

fn main() {
    let test_run = is_test_run();
    let start = Instant::now();

    giant_squid::run(test_run).unwrap();

    let duration = start.elapsed();

    println!("Execution time: {:?}", duration);
}

fn is_test_run() -> bool {
    let args: Vec<_> = env::args().collect();

    if args.len() < 2 { false } else { args[1] == "test" }
}
