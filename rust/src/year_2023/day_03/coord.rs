pub struct Coord {
    pub x: i32,
    pub y: i32,
}

impl Coord {
    pub fn new(x: i32, y: i32) -> Coord {
        Coord { x, y }
    }

    pub fn is_adjacent(&self, other: &Coord) -> bool {
        return (other.y == (self.y - 1)) && (other.x >= (self.x - 1) && other.x <= (self.x + 1))
        || (other.y == (self.y + 1)) && (other.x >= (self.x - 1) && other.x <= (self.x + 1))
        || (other.y == self.y) && (other.x == self.x - 1 || other.x == self.x + 1);
    }
}