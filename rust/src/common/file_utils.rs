use std::error::Error;
use std::fs::File;
use std::io::{BufRead, BufReader, Seek};
use std::path::PathBuf;

pub fn get_file_reader(year: &str, day: &str, test: bool) -> Result<(BufReader<File>, u32), Box<dyn Error>> {
    let path = PathBuf::from(format!(
        "src/year_{}/day_{}/{}.txt",
        year,
        day,
        if test { "test_input" } else { "input" },
    ));

    let mut file = File::open(path.as_path())?;

    let reader = BufReader::new(&file);
    let lines_count = reader.lines().count() as u32;

    // Reset the file pointer to the beginning of the file.
    file.seek(std::io::SeekFrom::Start(0))?;

    Ok((BufReader::new(file), lines_count))
}
