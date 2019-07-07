package computer

//Spec is exported struct, field model is not exported
type Spec struct { //exported struct  
    Maker string //exported field
    model string //unexported field
    Price int //exported field
}