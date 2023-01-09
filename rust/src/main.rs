mod common;
mod year_2021;

use std::time::Instant;
use year_2021::day_01::sonar_sweep;

fn main() {
    let start = Instant::now();

    sonar_sweep::run().unwrap();

    let duration = start.elapsed();

    println!("Execution time: {:?}", duration);
}
