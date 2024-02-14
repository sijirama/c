fn main() {
    let southern_germany = "Grüß Gott!";
    let japan = "ハロー・ワールド";

    let _regions = [southern_germany, japan];

    // for region in regions.iter() {
    //     println!("{}", region);
    // }

    println!("{}", &japan);
    calc();
    stuff();
}

fn calc() {
    let a = 10;
    let b: i32 = 20;

    let _c = add(a, b);
    //println!("a + b = {}", c);
}

fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn stuff() {
    let needle = 42;
    let haystack = [1, 1, 2, 3, 4, needle, 22, 32, 56, 545];
    for reference in haystack {
        let result = match reference {
            42 | 56 | 545 => "hit",
            _ => "miss",
        };

        if result == "hit" {
            println!("{}: {}", reference, result);
        }
    }
}
