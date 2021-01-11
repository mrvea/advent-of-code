package vacation82

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	day           = "8"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

type Instruction struct {
	Name      string
	Direction int
	Value     int
	Count     uint
}

func (I *Instruction) String() string {
	return fmt.Sprintf("Name: %s, direction: %d, value: %d, count: %d\n", I.Name, I.Direction, I.Value, I.Count)
}

func NewInstruction(name string, direction, value int) *Instruction {
	c := Instruction{
		Name:      name,
		Direction: direction,
		Value:     value,
		Count:     0,
	}
	return &c
}

/**

**/
func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".1 starting exectution...")
	filePath := basePath + inputName
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		filePath = basePath + testInputName
	}
	fmt.Printf("filePath: %s\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total := 0
	insts := make([]*Instruction, 0)
	for s.Scan() {
		fmt.Println(s.Text())
		n, d, v := processInstruction(s.Text())
		// bagName, contains := processDirections(s.Text())
		insts = append(insts, NewInstruction(n, d, v))

	}
	fmt.Println(insts)
PARENT:
	for sub, alt := range map[string]string{"nop": "jmp", "jmp": "nop"} {
		// program := make([]*Instruction, 0)
		for _, inst := range insts {
			if sub == inst.Name {
				inst.Name = alt

				// fmt.Println("new instructions")
				// fmt.Println(insts)
				total, err = execute(insts)
				if err == nil {
					fmt.Println(err)
					break PARENT
				}
				inst.Name = sub
				for _, inst = range insts {
					inst.Count = 0
				}
			}
		}
	}
	fmt.Println("total: ", total)
}

func execute(insts []*Instruction) (int, error) {
	i := 0
	total := 0
	for {
		if len(insts) <= i {
			fmt.Println("exected correctly")
			break
		}
		inst := insts[i]
		// fmt.Println(inst)
		if inst.Count > 0 {
			return 0, fmt.Errorf("loop")
		}
		switch inst.Name {
		case "nop":
			i++
		case "jmp":
			i += inst.Value * inst.Direction
		case "acc":
			i++
			total += inst.Value * inst.Direction
		default:
			log.Panicf("unknown instruction %s", inst.Name)
		}
		inst.Count++
	}

	return total, nil
}

func processInstruction(raw string) (string, int, int) {
	var n, dStr string
	var d int
	fmt.Sscanf(raw, "%s %s", &n, &dStr)
	switch dStr[0] {
	case '+':
		d = 1
	case '-':
		d = -1
	default:
		log.Panicf("unexpected %s direction string", dStr)
	}
	v, err := strconv.Atoi(dStr[1:])
	if err != nil {
		log.Panic("value is not a number")
	}
	return n, d, v
}
