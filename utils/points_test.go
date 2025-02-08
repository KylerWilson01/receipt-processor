package utils_test

import (
	"testing"

	"github.com/KylerWilson01/receipt-processor.git/models"
	"github.com/KylerWilson01/receipt-processor.git/utils"
)

func TestCheckRetailerName(t *testing.T) {
	tests := []struct {
		name, against string
		want          int
	}{
		{"All AlphaNumeric Characters", "Target", 6},
		{"Mix AlphaNumeric and Non AlphaNumeric Characters", "M&M Corner Market", 14},
		{"Empty", "", 0},
		{"Only Non-AlphaNumeric Characters", "&&&&&    &&&", 0},
	}

	pointUtil := utils.PointUtil{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := pointUtil.CheckRetailerName(test.against)

			if test.want != got {
				t.Fatalf("With %s we want %d but got %d", test.against, test.want, got)
			}
		})
	}
}

func TestCheckRoundDollar(t *testing.T) {
	tests := []struct {
		name, against string
		want          int
	}{
		{"Round Dollar", "9.00", 50},
		{"Contains cents", "9.09", 0},
		{"Zero", "0.00", 0},
		{"Negative", "-1.00", 0},
	}

	pointUtil := utils.PointUtil{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := pointUtil.CheckRoundDollar(test.against)

			if test.want != got {
				t.Fatalf("With %s we want %d but got %d", test.against, test.want, got)
			}
		})
	}
}

func TestCheckMultipleOfQuarter(t *testing.T) {
	tests := []struct {
		name, against string
		want          int
	}{
		{"Round Dollar", "9.00", 25},
		{"Contains 9 cents", "9.09", 0},
		{"Contains 25 cents", "9.25", 25},
		{"Contains 50 cents", "9.50", 25},
		{"Contains 75 cents", "9.75", 25},
		{"Contains 78 cents", "9.78", 0},
		{"Zero", "0.00", 0},
		{"Negative", "-1.00", 0},
	}

	pointUtil := utils.PointUtil{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := pointUtil.CheckMultiple(test.against)

			if test.want != got {
				t.Fatalf("With %s we want %d but got %d", test.against, test.want, got)
			}
		})
	}
}

func TestCountLengthOfItems(t *testing.T) {
	tests := []struct {
		name    string
		against []models.Item
		want    int
	}{
		{"1 item", []models.Item{{}}, 0},
		{"2 items", []models.Item{{}, {}}, 5},
		{"0 items", []models.Item{}, 0},
		{"3 items", []models.Item{{}, {}, {}}, 5},
		{"8 items", []models.Item{{}, {}, {}, {}, {}, {}, {}, {}}, 20},
	}

	pointUtil := utils.PointUtil{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := pointUtil.CountLengthOfItems(test.against)

			if test.want != got {
				t.Fatalf("With %s we want %d but got %d", test.against, test.want, got)
			}
		})
	}
}

func TestDescriptionLengthOfItem(t *testing.T) {
	tests := []struct {
		name    string
		against models.Item
		want    int
	}{
		{
			"Not Multiple of three",
			models.Item{ShortDescription: "Mountain Dew 12PK", Price: "12.25"},
			0,
		},
		{"Length is 18", models.Item{ShortDescription: "Emils Cheese Pizza", Price: "12.25"}, 3},
		{
			"The description has white space around",
			models.Item{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12"},
			3,
		},
		{
			"Description is white space with a character",
			models.Item{ShortDescription: "     A      ", Price: "12.25"},
			0,
		},
	}

	pointUtil := utils.PointUtil{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := pointUtil.CheckDescriptionLength(test.against)

			if test.want != got {
				t.Fatalf("With %s we want %d but got %d", test.against, test.want, got)
			}
		})
	}
}

func TestCheckDate(t *testing.T) {
	tests := []struct {
		name, against string
		want          int
	}{
		{"Invalid Date", "2000-02-30", 0},
		{"Valid Date with points", "2000-02-03", 6},
		{"Valid Date with no points", "2000-02-04", 0},
	}

	pointUtil := utils.PointUtil{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := pointUtil.CheckDate(test.against)

			if test.want != got {
				t.Fatalf("With %s we want %d but got %d", test.against, test.want, got)
			}
		})
	}
}

func TestCheckTime(t *testing.T) {
	tests := []struct {
		name, against string
		want          int
	}{
		{"Invalid time", "24:05", 0},
		{"Valid Date with points", "14:33", 10},
		{"Valid Date with no points on 2pm", "14:00", 0},
		{"Valid Date with no points on 4pm", "16:00", 0},
		{"Valid Date with no points after 4pm", "20:00", 0},
		{"Valid Date with no points before 2pm", "10:00", 0},
	}

	pointUtil := utils.PointUtil{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := pointUtil.CheckTime(test.against)

			if test.want != got {
				t.Fatalf("With %s we want %d but got %d", test.against, test.want, got)
			}
		})
	}
}
