mod common;
mod year_2021;

use std::time::Instant;
use crate::year_2021::day_02::dive;

fn main() {
    let start = Instant::now();

    dive::run().unwrap();

    let duration = start.elapsed();

    println!("Execution time: {:?}", duration);
}
