package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type customer struct {
	ID          int
	name        string
	street      string
	city        string
	state       string
	zip         int
	email       string
	phoneNumber int
}

func main() {
	c1 := customer{341, "Jane Smith", "101 E Main St", "Springfield", "MO", 38249, "janesmith@email.com", 1234567890}
	c2 := customer{1, "John Doe", "304 W Maple St", "Pittsburg", "PA", 88218, "john1980@email.com", 2223334444}
	c3 := customer{1020, "Carl Swanson", "1200 W El Padre Ave", "Jacksonville", "FL", 92843, "swanson1111@email.com", 9993332222}
	c4 := customer{22, "Wade Johnson", "3215 N Sante Fe Ave", "Albuquerque", "NM", 97253, "wjohnson1234@email.com", 4441118888}
	c5 := customer{34, "George Hamiliton", "10 S Clark", "Houston", "TX", 23853, "george2000@email.com", 8882221111}
	c6 := customer{200, "Beth Anderson", "202 E Apple Ln", "Seattle", "WA", 97234, "beth-anderson@email.com", 7772224444}
	customers := []customer{c1, c2, c3, c4, c5, c6}
	for _, v := range customers {
		fmt.Println(v.ID, v.name, v.street, v.city, v.street, v.zip, v.email, v.phoneNumber)
	}
	fmt.Println("")
	/*
		minwidth	minimal cell width including any padding
		tabwidth	width of tab characters (equivalent number of spaces)
		padding		padding added to a cell before computing its width
		padchar		ASCII char used for padding
					if padchar == '\t', the Writer will assume that the
					width of a '\t' in the formatted output is tabwidth,
					and cells are left-aligned independent of align_left
					(for correct-looking results, tabwidth must correspond
					to the tab width in the viewer displaying the result)
		flags		formatting control
	*/
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, '.', tabwriter.Debug)
	// w := tabwriter.NewWriter(os.Stdout, 10, 0, 0, '.', tabwriter.Debug)
	// w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, '.', tabwriter.Debug)
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, '.', tabwriter.Debug)
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, '.', tabwriter.Debug)
	// w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.TabIndent) // most systems use 4 or 8 width
	// w := tabwriter.NewWriter(os.Stdout, 0, 3, 1, '\t', tabwriter.TabIndent) // this is not same as viewer tab width and is purposely wrong
	// w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "ID\tName\tStreet\tCity\tState\tZIP\tEmail\tPhone Number\t")
	for _, v := range customers {
		fmt.Fprint(w, v.ID, "\t", v.name, "\t", v.street, "\t", v.city, "\t", v.state, "\t", v.zip, "\t", v.email, "\t", v.phoneNumber, "\t\n")
		// fmt.Fprint(w, v.ID, "	", v.name, "	", v.street, "	", v.city, "	", v.state, "	", v.zip, "	", v.email, "	", v.phoneNumber, "\t\n") // poor readibility
	}

	/*
		// https://golang.org/pkg/text/tabwriter/#example__elastic
		fmt.Fprintln(w, "aa\tbb\tc")
		fmt.Fprintln(w, "aa\tbb\tcc")
		fmt.Fprintln(w, "aaa\t") // trailing tab
		fmt.Fprintln(w, "aaaa\tdddd\teeee")
	*/

	// func (b *Writer) Flush() error
	w.Flush()
}
