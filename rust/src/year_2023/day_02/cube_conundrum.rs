use std::error::Error;
use std::io::BufRead;

use crate::common::file_utils::get_file_reader;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "02", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    Ok(())
}