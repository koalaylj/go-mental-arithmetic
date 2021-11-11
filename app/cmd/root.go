package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/koalaylj/go-mental-arithmetic/app/pkg/op"
	"github.com/koalaylj/go-mental-arithmetic/app/pkg/pdf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "hugo",
		Short: "四则运算生成器",
		Long:  `小学生的梦魇，四则运算生成器。可配置10以内20以内以及无限制的加减运算，可设置是否进位，是否包括乘除等等`,
		Run: func(cmd *cobra.Command, args []string) {
			m := pdf.New()

			pdf.SetHeader(m, arg_row, arg_col)
			pdf.SetFooter(m)

			for i := 0; i < arg_pages; i++ {
				cells := innerText()
				pdf.BuildPage(m, cells)
				if i < arg_pages-1 {
					m.AddPage()
				}
			}

			pdf.Export(m, arg_path)
		},
	}
)

var (
	arg_min   = 1
	arg_max   = 10
	arg_carry = false
	arg_pages = 10
	arg_path  = "./pdf"
	arg_col   = 4
	arg_row   = 15
	ops       = []string{}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile, "file", "f", "", "config")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())

		arg_add := viper.GetBool("op.add")
		arg_sub := viper.GetBool("op.sub")
		arg_mul := viper.GetBool("op.mul")
		arg_div := viper.GetBool("op.div")

		if arg_add {
			ops = append(ops, "+")
		}

		if arg_sub {
			ops = append(ops, "-")
		}

		if arg_mul {
			ops = append(ops, "*")
		}

		if arg_div {
			ops = append(ops, "/")
		}

		if len(ops) == 0 {
			panic("至少指定一种运算 add sub mul div")
		}

		arg_min = viper.GetInt("difficulty.min")
		arg_max = viper.GetInt("difficulty.max")
		arg_carry = viper.GetBool("difficulty.carry")
		arg_path = viper.GetString("pdf.path")
		arg_pages = viper.GetInt("pdf.pages")
		arg_col = viper.GetInt("pdf.col")
		arg_row = viper.GetInt("pdf.row")

	} else {
		fmt.Println("initConfig failed", err)
	}
}

func randomOp(random *rand.Rand) string {
	index := random.Intn(len(ops))
	return ops[index]
}

func getRandom() *rand.Rand {
	now := time.Now().UnixNano()
	seed := rand.NewSource(now)
	random := rand.New(seed)
	return random
}

func innerText() [][]string {
	cells := [][]string{}

	for row_i := 0; row_i < arg_row; row_i++ {

		row := []string{}

		for col_i := 0; col_i < arg_col; col_i++ {
			no := row_i + col_i*arg_row + 1

			random := getRandom()

			operand := randomOp(random)

			item := ""

			switch operand {
			case "+":
				item = op.RandomAdd(random, arg_max, arg_min)
			case "-":
				item = op.RandomSub(random, arg_max, arg_min)
			}

			cell := fmt.Sprintf("%2d) %s\n", no, item)

			row = append(row, cell)
		}

		cells = append(cells, row)
	}

	return cells
}
