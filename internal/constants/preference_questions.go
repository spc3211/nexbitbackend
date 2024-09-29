package constants

var Questions = []struct {
	ID int
    Question string
    Answers  []struct {
		Id    int
        Text  string
        Score int
    }
}{
    {
		ID: 1,
        Question: "What is your primary investment goal?",
        Answers: []struct {
			Id    int
            Text  string
            Score int
        }{
            {1, "Preserve my wealth and earn stable returns", 1},
            {2, "Grow my wealth steadily over time", 2},
            {3, "Maximize my wealth growth potential", 3},
        },
    },
    {
		ID: 2,
        Question: "How long do you plan to keep your money invested?",
        Answers: []struct {
			Id    int
            Text  string
            Score int
        }{
            {1, "1-3 years", 1},
            {2, "4-7 years", 2},
            {3, "8+ years", 3},
        },
    },
    {
		ID: 3,
        Question: "Which statement best describes your investment knowledge?",
		Answers: []struct {
			Id    int
			Text  string
			Score int
		}{
			{1, "I'm new to investing", 1},
            {2, "I have some investing experience", 2},
            {3, "I'm an experienced investor", 3},	
		},
    },
	{
		ID: 4,
        Question: "How much of your total savings are you planning to invest?",
		Answers: []struct {
			Id    int
			Text  string
			Score int
		}{
			{1, "Less than 25%", 1},
            {2, "25% to 50%", 2},
            {3, "More than 50%", 3},
		},
    },
}
