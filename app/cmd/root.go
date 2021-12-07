package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/koalaylj/go-mental-arithmetic/app/pkg/config"
	"github.com/koalaylj/go-mental-arithmetic/app/pkg/op"
	"github.com/koalaylj/go-mental-arithmetic/app/pkg/pdf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	arg_col = 4
	arg_row = 15

	options = struct {
		add   config.OP_ADD
		sub   config.OP_SUB
		pages int
		path  string
		ops   []string
	}{
		add:   config.OP_ADD{},
		sub:   config.OP_SUB{},
		pages: 10,
		path:  "./",
		ops:   []string{},
	}
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

			for i := 0; i < options.pages; i++ {
				cells := innerText()
				pdf.BuildPage(m, cells)
				if i < options.pages-1 {
					m.AddPage()
				}
			}

			pdf.Export(m, options.path)
		},
	}
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

		options.ops = viper.GetStringSlice("ops")

		if len(options.ops) == 0 {
			panic("至少指定一种运算 + - ")
		}

		options.add.Min = viper.GetInt("add.min")
		options.add.Max = viper.GetInt("add.max")
		options.add.Carry = viper.GetBool("add.carry")
		options.add.Bounds = viper.GetIntSlice("add.bounds")

		options.sub.Min = viper.GetInt("sub.min")
		options.sub.Max = viper.GetInt("sub.max")
		options.sub.Borrow = viper.GetBool("sub.borrow")
		options.sub.Bounds = viper.GetIntSlice("sub.bounds")

		options.path = viper.GetString("pdf.path")
		options.pages = viper.GetInt("pdf.pages")

		fmt.Printf("%+v", options)

	} else {
		fmt.Println("initConfig failed", err)
	}
}

func randomOp(random *rand.Rand) string {
	index := random.Intn(len(options.ops))
	return options.ops[index]
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
				item = op.RandomAdd(random, options.add)
			case "-":
				item = op.RandomSub(random, options.sub)
			}

			cell := fmt.Sprintf("%2d) %s\n", no, item)

			row = append(row, cell)
		}

		cells = append(cells, row)
	}

	return cells
}
