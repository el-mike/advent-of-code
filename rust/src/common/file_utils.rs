use std::error::Error;
use std::fs::File;
use std::io::BufReader;

pub fn get_file_reader(day: &str, test: bool) -> Result<BufReader<File>, Box<dyn Error>> {
    let file_name = if test { "test_input" } else { "input" };
    let file = File::open(format!("./src/year_2021/day_{}/{}.txt", day, file_name))?;

    return Ok(BufReader::new(file));
}
