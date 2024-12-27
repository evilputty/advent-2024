use std::fs::File;
use std::io::Read;

fn main() {
    let mut file = match File::open("input_data.txt") {
        Ok(file) => file,
        Err(e) => panic!("Unable to open file: {:?}", e),
    };

    let mut contents = String::new();

    file.read_to_string(&mut contents).expect("Unable to read to string");

    let mut left: Vec<i32> = Vec::new();
    let mut right: Vec<i32> = Vec::new();

    for value in contents.lines()   {
        let mut count: u8 = 0;
        for x in value.split("   "){
            if count <= 0 {
                left.push(x.parse().expect("cant convert"));
                count += 1;
            } else {
                right.push(x.parse().expect("cant convert"));
            }
        }
    }
    
    println!("{:?}", left);

}
