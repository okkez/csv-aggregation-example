use std::collections::BTreeMap;
use std::env;
use std::error::Error;

use serde::Deserialize;

#[derive(Deserialize, Debug)]
struct Record {
    id: u64,
    name: String,
    description: String,
    cost: f64,
}

fn run() -> Result<(), Box<dyn Error>> {
    let path = env::args().nth(1).unwrap();
    let mut csv_reader = csv::ReaderBuilder::new()
        .has_headers(true)
        .from_path(&path)?;
    let mut name_to_cost = BTreeMap::new();

    for result in csv_reader.deserialize() {
        let record: Record = result?;
        let cost = record.cost;
        *name_to_cost.entry(record.name).or_insert(0.0) += cost;
    }

    for (name, cost) in name_to_cost {
        println!("{}\t{:.3}", name, cost);
    }

    Ok(())
}

fn main() {
    run().unwrap()
}
