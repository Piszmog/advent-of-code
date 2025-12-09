use std::env;
use std::error::Error;
use std::fs::File;
use std::io::BufReader;
use std::time::Instant;

pub fn get_args() -> Vec<String> {
    env::args().collect()
}

pub fn get_csv_reader(file_path: &str, has_headers: bool, delimiter: u8) -> Result<csv::Reader<std::fs::File>, Box<dyn Error>> {
    let file = std::fs::File::open(file_path)?;
    let reader = csv::ReaderBuilder::new()
        .has_headers(has_headers)
        .delimiter(delimiter)
        .from_reader(file);
    Ok(reader)
}

pub fn get_file_reader(file_path: &str) -> Result<BufReader<File>, std::io::Error> {
    File::open(file_path)
        .and_then(|f| Ok(BufReader::new(f)))
}

pub fn print_elapsed_time(start: Instant) {
    let duration = Instant::now().duration_since(start);
    println!("Duration: {:?}", duration);
}