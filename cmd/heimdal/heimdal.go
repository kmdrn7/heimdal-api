package heimdal

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "heimdal",
	Short: "Heimdal is a tools for communicating with Biforst API",
	Long:  "API based CLI for communicating with Bifrost API.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
''......--------::::::::::::::::::-------........'
''.......--------::://:::::::+o:::-------......'''
'''......------:::+s/:::::::::oo/-------......''''
''''......-----:-:/o/:--//::/oyy+/:----......'''''
'''......------------::sdds++/:/:-----........''''
'''......------------:/:o++:::::-------..........'
''......------------o/+://+/o-::--------.........'
'.......-----:::::::o+o:o+++:+:::::-----..........
'......---:/+syyys+s::+yhdy/:+ooyyo+/:---.........
'.....----:://////-:::/+oho/:/:++/:/:/:---:.......
'.....:--..-------...:++osss:--:-:::::---:/--.....
''''''.......---+y/+osysdmyyys+oh---......''''''''
		`)
		fmt.Println("I am Heimdal, the guardian of Bifrost ............")
		fmt.Println("I see everything, I hear everything ..............")
		fmt.Println()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
