use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2021", "05", test_run)
        .unwrap_or_else(|err| panic!("{}", err));

    for line in reader.lines() {

    }

    Ok(())
}
