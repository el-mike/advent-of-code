use std::error::Error;
use std::fs::File;
use std::io::BufReader;
use std::path::PathBuf;

pub fn get_file_reader(day: &str, test: bool) -> Result<BufReader<File>, Box<dyn Error>> {
    let path = PathBuf::from(format!(
        "src/year_2021/day_{}/{}.txt",
        day,
        if test { "test_input" } else { "input" },
    ));

    let file = File::open(path.as_path())?;

    Ok(BufReader::new(file))
}
