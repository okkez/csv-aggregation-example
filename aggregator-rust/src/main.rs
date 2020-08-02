use std::collections::HashMap;
use std::env;
use std::error::Error;
use std::fs::File;
use std::io::BufReader;

use serde::{Deserialize};

#[derive(Deserialize, Debug)]
struct Record {
    id: u64,
    name: String,
    description: String,
    cost: f64,
}

fn run() -> Result<(), Box<dyn Error>> {
    let args: Vec<String> = env::args().collect();
    let f = File::open(&args[1])?;
    let b = BufReader::new(f);
    let mut csv_reader = csv::ReaderBuilder::new().has_headers(true).from_reader(b);
    let mut name_to_cost = HashMap::new();

    for result in csv_reader.deserialize() {
        let record: Record = result?;
        let cost = match name_to_cost.get(&record.name) {
            Some(value) => *value,
            None => 0.0,
        };
        name_to_cost.insert(record.name.clone(), record.cost + cost);
    }


    let mut v = name_to_cost.into_iter().collect::<Vec<_>>();
    v.sort_by(|x, y| x.0.cmp(&y.0));
    for (name, cost) in v {
        println!("{}\t{:.3}", name, cost);
    }
    
    Ok(())
}

fn main() {
    run().unwrap()
}
