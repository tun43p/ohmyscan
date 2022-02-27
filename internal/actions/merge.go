package actions

import (
	"github.com/urfave/cli/v2"
	"github.com/tun43p/ohmyscan/internal/utils"
)

// Merge ... Function that allows us to merge the downloaded scans according to the given parameters.
// 	- c(*cli.Content): Pass the application context.
func Merge(c *cli.Context) {
	name := c.String("name")
	number := c.String("number")

	if name == "" || number == "" {
		utils.Message(utils.ErrorArgumentsEmpty, "error")
	}

	utils.ConvertToPDF(name, number)
}
