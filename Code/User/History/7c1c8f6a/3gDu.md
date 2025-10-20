# arg parser
if len(os.Args) > 1 {
		fmt.Println("First Arguement: ", os.Args[1])
	}


	
func printBanner() {
	color.Red(`
   ____  __  ____   ______
  / __ \/  |/  / | / /  _/
 / / / / /|_/ /  |/ // /
/ /_/ / /  / / /|  // /
\____/_/  /_/_/ |_/___/

	`)
}


cmd := exec.Command("foo", "bar")
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr
err := cmd.Run()
if err != nil {
    fmt.Println("Error:", err)
}


arguements:
0: file name
1: first arguement
2: second arguement
3..


get current program name
```
import(
	"os"
	"fmt"
	"path/filepath"
)
func main() {
	progPath := os.Args[0]
	progName := filepath.Base(progPath)
	fmt.Println("%s", progName)
}

