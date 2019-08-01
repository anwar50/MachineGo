// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"time"
// )

// type PageRadioButtons struct {
// 	Name       string
// 	Value      string
// 	isDisabled bool
// 	IsChecked  bool
// 	Text       string
// }
// type PageVariables struct {
// 	Date string
// 	Time string
// }
// type Todo struct {
// 	Title string
// 	Done  bool
// }

// type TodoPageData struct {
// 	PageTitle string
// 	Projects     []Todo
// }

// func hello_world(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }
// func Navbar(w http.ResponseWriter, r *http.Request) {

// }
// func main() {
// 	// http.HandleFunc("/", Home)

// 	tmpl := template.Must(template.ParseFiles("layout.html"))

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		data := TodoPageData{
// 			PageTitle: "My TODO list",
// 			Projects: []Todo{
// 				{Title: "Task 1", Done: false},
// 				{Title: "Task 2", Done: true},
// 				{Title: "Task 3", Done: true},
// 			},
// 		}
// 		tmpl.Execute(w, data)
// 	})

// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

// func Home(w http.ResponseWriter, r *http.Request) {
// 	now := time.Now()
// 	Home_Variables := PageVariables{
// 		Date: now.Format("02-01-2006"),
// 		Time: now.Format("15-04-15"),
// 	}
// 	//parse the html file to the index
// 	t, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		log.Print("template parsing error: ", err)
// 	}
// 	err = t.Execute(w, Home_Variables)
// 	if err != nil {
// 		//log the error if there is one
// 		log.Print("template executing error: ", err)
// 	}

// }

package main

import (
	"html/template"
	"net/http"
	"time"
)

type PageRadioButtons struct {
	Name       string
	Value      string
	isDisabled bool
	IsChecked  bool
	Text       string
}

type Python_Project struct {
	Description string
	Problem     string
	Solution    string
	Done        bool
}
type JavaScript_Project struct {
	Description string
	Problem     string
	Solution    string
	Done        bool
}
type MachineLearning_Description1 struct {
	SupervisedLearning string
	Done               bool
}
type MachineLearning_Description2 struct {
	UnSupervisedLearning string
	Done                 bool
}
type MachineLearning_Description3 struct {
	ReinforcementLearning string
	Done                  bool
}
type TodoPageData struct {
	PageTitle                     string
	ColourTitle                   string
	Machine                       string
	What_Is_ML                    string
	Supervised                    string
	UnSupervised                  string
	Reinforcement                 string
	Python_Projects               []Python_Project
	JavaScript_Projects           []JavaScript_Project
	MachineLearning_Supervised    []MachineLearning_Description1
	MachineLearning_UnSupervised  []MachineLearning_Description2
	MachineLearning_Reinforcement []MachineLearning_Description3
	Date                          string
	Time                          string
}

func main() {

	tmpl := template.Must(template.ParseFiles("layout2.html"))
	now := time.Now()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle:     "MachineGo Projects",
			Machine:       "What is machine learning?",
			ColourTitle:   "Please change colour to see whether it will be predicted light or dark! watch the time :)",
			What_Is_ML:    "Machine learning uses a variety of algorithms that iteratively learn from data to improve, describe data, and predict outcomes.",
			Supervised:    "Supervised Learning",
			UnSupervised:  "Unsupervised Learning",
			Reinforcement: "Reinforcement Learning",
			Python_Projects: []Python_Project{
				{Description: "Task 1", Problem: "sfd", Solution: "dgs", Done: false},
			},
			JavaScript_Projects: []JavaScript_Project{
				{Description: "Task 56", Problem: "dsgkfhj", Solution: "kdsuh", Done: true},
			},
			MachineLearning_Supervised: []MachineLearning_Description1{
				{SupervisedLearning: "In supervised learning we start with a dataset that has training examples, each example has an assiocated label which identifies it.", Done: true},
			},
			MachineLearning_UnSupervised: []MachineLearning_Description2{
				{UnSupervisedLearning: "Unsupervised learning is quite different from supervised in the sense that it almost always does not have a definite output. The learning agent aims to find structures or patterns in the data.", Done: true},
			},
			MachineLearning_Reinforcement: []MachineLearning_Description3{
				{ReinforcementLearning: "Reinforcement learning is where the learner receives rewards and punishments for its actions. The reward could simply be utility and the agent could be told to receive as much utility as possible in order to “win”. Utility here could just be a normal variable.", Done: true},
			},
			Date: now.Format("02-01-2006"),
			Time: now.Format("15-04-15"),
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8000", nil)
}
