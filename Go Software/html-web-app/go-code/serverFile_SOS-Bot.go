// # sos-bot

package main


// . ° ~ +
import (
    "os"
    "log"

        "text/template"
        "net/http"

        "time"

        "cloud.google.com/go/firestore"

        firebase "firebase.google.com/go"
  "google.golang.org/api/option"


)

// . ° ~ +
type classSchedule struct {
    weekDay string
    dayTime string
    month string
    year string
    className string
    classID string

}

type studentProfiles struct {
    studentName string
    studentID string

}

type paymentInformation struct {
    studentName string
    studentID string
    statusDate string
    paymentStatus string
    paymentHistory []string

}


type studentAttendence struct {


}

type classEvent struct {


}

type studentHomework struct {
    studentName string
    studentID string


}
// . ° ~ +
type studentFeildtrips struct {


}

type classQuestions struct {


}

type classMovies struct {


}

type studentWorkshops struct {


}

// . ° ~ +
var classSchedule[] {

    weekDays : ['Mon','Wed','Fri'],
    timeDay : ['10:am','2:pm']


}


// . ° ~ +
func addClassToList(ctx context.Context, client *firestore.Client, addClass) error {
  _ClassListData := addClass

  _, err := client.Collection("ClassList").Doc("ClassList").Set(ctx, _ClassListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}




// . ° ~ +
func listClass() {


}

func loadClass() {


}

func joinClass() {


}

func leaveClass() {


}

// . ° ~ +
func addDocWithID(ctx context.Context, client *firestore.Client) error {
  var data = make(map[string]interface{})

  _, err := client.Collection("class").Doc("new-class-id").Set(ctx, data)


  if err != nil {
     log.Printf("An error has occurred: %s", err)
  }
  return err
}
// . ° ~ +
func addDocDataTypes(ctx context.Context, client *firestore.Client) error {
        doc := make(map[string]interface{})
        doc["stringExample"] = "Hello world!"
        doc["booleanExample"] = true
        doc["numberExample"] = 3.14159265
        doc["dateExample"] = time.Now()
        doc["arrayExample"] = []interface{}{5, true, "hello"}
        doc["nullExample"] = nil
        doc["objectExample"] = map[string]interface{}{
                "a": 5,
                "b": true,
        }

  _, err := client.Collection("data").Doc("one").Set(ctx, doc)

  if err != nil {
   // Handle any errors
    log.Printf("An error has occurred: %s", err)
 }
  return err
}
// . ° ~ +
func addDocAsEntity(ctx context.Context, client *firestore.Client) error {
  studentData := studentData{
    Name:    "Phoenix",
    Status: "Active",
}

  _, err := client.Collection("studentData").Doc("studentData").Set(ctx, studentData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}


// . ° ~ +
func addHomeworkToFile(ctx context.Context, client *firestore.Client, saveHomework) error {
  homeworkData := saveHomework

  _, err := client.Collection("HomeworkFile").Doc("studentData").Set(ctx, homeworkData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// . ° ~ +
func addStudentAttendence(ctx context.Context, client *firestore.Client, attendenceData) error {
  studentData := attendenceData

  _, err := client.Collection("AttendenceData").Doc("studentAttendence").Set(ctx, studentData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}




// . ° ~ +
func startClass() {
    var user := studentProfiles[] {}

    var class := classSchedule[] {}

    var here := studentAttendence[] {}

    var classTimer := ':'

    addStudentAttendece(attendenceData)

}


// . ° ~ +
func meetTeacher() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}


}



// . ° ~ +
func addEventToList(ctx context.Context, client *firestore.Client, addEvent) error {
  _EventListData := addEvent

  _, err := client.Collection("EventList").Doc("EventList").Set(ctx, _EventListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}
// . ° ~ +
func doEvent() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}

    var event := eventData[] {}

}


// . ° ~ +
func addHomeworkToList(ctx context.Context, client *firestore.Client, addFeildtrip) error {
  _FeildtripListData := addFeildtrip

  _, err := client.Collection("FeildtripList").Doc("FeildtripList").Set(ctx, _FeildtripListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// . ° ~ +
func homework() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}
// receives homework files from Discord User Interface
    var saveHomework := homeworkData[] {}
    action := getActionType() {
    }
    if action == turnIn {
        addHomeworkToFile(saveHomework)
    }
}

// . ° ~ +
func addQuestionToList(ctx context.Context, client *firestore.Client, addFeildtrip) error {
  _FeildtripListData := addFeildtrip

  _, err := client.Collection("FeildtripList").Doc("FeildtripList").Set(ctx, _FeildtripListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// . ° ~ +
func askQuestions() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}

    var question := classQuestions[] {}

}
// . ° ~ +
func meetStudents() {

    var class := classSchedule[] {}

    var here := studentAttendence[] {}


}


// . ° ~ +
func addFeildtripToList(ctx context.Context, client *firestore.Client, addFeildtrip) error {
  _FeildtripListData := addFeildtrip

  _, err := client.Collection("FeildtripList").Doc("FeildtripList").Set(ctx, _FeildtripListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}


// . ° ~ +
func studentFeildtrip() {

    var class := classSchedule[] {}

    var here := studentAttendence[] {}

    var trip := studentFeildtrips[] {}

}



// . ° ~ +
func addMovieToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}


// . ° ~ +
func watchMovie() {
    var class := classSchedule[] {}
    var here := studentAttendence[] {}

}


// ~ addNewWorkehop
func addWorkshopToList(ctx context.Context, client *firestore.Client, addWorkshop) error {
  _WorkshopListData := addWorkshop

  _, err := client.Collection("WorkshopList").Doc("WorkshopList").Set(ctx, _WorkshopListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// . ° ~ +
func studentWorkshops() {


}


// ~ addNewTeacher
func addTeacherToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// ~ addNewStudent
func addStudentToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// ~ addNewMentor
func addMentorToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// ~ addNewStaff
func addStaffToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// ~ addNewProject
func addProjectToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// . ° ~ +
func openProject() {
    var class := classSchedule[] {}
    var here := studentAttendence[] {}

}



// ~ addNewSpeaker
func addSpeakerToList(ctx context.Context, client *firestore.Client, addSpeaker) error {
  _MovieListData := addMovie

  _, err := client.Collection("SpeakerList").Doc("SpeakerList").Set(ctx, _SpeakerListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// ~ addNewLecture
func addLectureToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)

  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}

// . ° ~ +
func attendLecture() {
    var class := classSchedule[] {}
    var here := studentAttendence[] {}

}


// . ° ~ +
func endClass() {
    var user := studentProfiles[] {}

    var class := classSchedule[] {}

    var here := studentAttendence[] {}

}

// ~ addNew, Function 
func addNew(w http.ResponseWriter, r *http.Request) {


// ~ addNewTeacher
    if r.URL.Path == "addNewTeacher" {
    	addTeacherToList();
    }

// ~ addNewStudent
    if r.URL.Path == "addNewStudent" {
    	addStudentToList();
    }

// ~ addNewMentor
    if r.URL.Path == "addNewMentor" {
    	addMentorToList();
    }

// ~ addNewStaff
    if r.URL.Path == "addNewStaff" {
    	addStaffToList();
    }

// ~ addNewHomework
    if r.URL.Path == "addNewHomework" {
    	addHomeworkToList();
    }

// ~ addNewProject
    if r.URL.Path == "addNewProject" {
    	addProjectToList();
    }

// ~ addNewMovie
    if r.URL.Path == "addNewMovie" {
    	addMovieToList();
    }

// ~ addNewFeildtrip
    if r.URL.Path == "addNewFeildtrip" {
    	addFeildtripToList();
    }

// ~ addNewWorkshop
    if r.URL.Path == "addNewWorkshop" {
    	addWorkshopToList();
    }

// ~ addNewSpeaker
    if r.URL.Path == "addNewSpeaker" {
    	addSpeakerToList();
    }

// ~ addNewLecture
    if r.URL.Path == "addNewLecture" {
    	addLectureToList();
    }

// ~ addNewQuestion
    if r.URL.Path == "addNewQuestion" {
    	addQuestionToList();
    }


} // END ~ addNew func


// ~ getDataList, Function 
func addNew(w http.ResponseWriter, r *http.Request) {


// ~ getDataList_Teacher
    if r.URL.Path == "getDataList_Teacher" {
    	addTeacherToList();
    }

// ~ getDataList_Student
    if r.URL.Path == "getDataList_Student" {
    	addStudentToList();
    }

// ~ getDataList_Mentor
    if r.URL.Path == "getDataList_Mentor" {
    	addMentorToList();
    }

// ~ getDataList_Staff
    if r.URL.Path == "getDataList_Staff" {
    	addStaffToList();
    }

// ~ getDataList_Homework
    if r.URL.Path == "getDataList_Homework" {
    	addHomeworkToList();
    }

// ~ getDataList_Project
    if r.URL.Path == "getDataList_Project" {
    	addProjectToList();
    }

// ~ getDataList_Movie
    if r.URL.Path == "getDataList_Movie" {
    	addMovieToList();
    }

// ~ getDataList_Feildtrip
    if r.URL.Path == "getDataList_Feildtrip" {
    	addFeildtripToList();
    }

// ~ getDataList_Workshop
    if r.URL.Path == "getDataList_Workshop" {
    	addWorkshopToList();
    }

// ~ getDataList_Speaker
    if r.URL.Path == "getDataList_Speaker" {
    	addSpeakerToList();
    }

// ~ getDataList_Lecture
    if r.URL.Path == "getDataList_Lecture" {
    	addLectureToList();
    }

// ~ getDataList_Question
    if r.URL.Path == "getDataList_Question" {
    	addQuestionToList();
    }


} // END ~ getDataList func


// . ° ~ +
// . indexHandler
func indexHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
    }

// ,

// . ° ~ +
  pageTitle := "SOS Bot App"
  pagePath := r.URL.Path
// . ° ~ +
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,

      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },

  }  //. .  pageData

// . ° ~ +
  pageFilePath := template.Must(template.ParseFiles("main_layout.html"))
  pageFilePath.Execute(w, pageData)

}  //  .  indexHandler



func main() {

  appName := "SOS Bot App"
// . ° ~ +
  http.HandleFunc("/", indexHandler)
// . ° ~ +
  http.HandleFunc("/action/about", aboutBot)
  
// . ° ~ +
  http.HandleFunc("/action/listclass", listClass)
  http.HandleFunc("/action/loadclass", loadClass)
  http.HandleFunc("/action/joinclass", joinClass)
  http.HandleFunc("/action/leaveclass", leaveClass)

// . ° ~ +
 http.HandleFunc("/action/classSchedule", classSchedule)

// . ° ~ +
 http.HandleFunc("/action/startclass", startClass)
 http.HandleFunc("/action/endclass", endClass)
 
 // . ° ~ +
  http.HandleFunc("/action/meetteacher", meetTeacher)
  http.HandleFunc("/action/meetstudents", meetStudents)
 
 // . ° ~ +
  http.HandleFunc("/action/askquestions", askQuestions)
 
// . ° ~ +
 http.HandleFunc("/action/veiwprojects", veiwProjects)
 http.HandleFunc("/action/homework", homework)

// . ° ~ +
 http.HandleFunc("/action/doevent", doEvents)
 
// . ° ~ +
 http.HandleFunc("/action/watchmovie", watchMovie)
 http.HandleFunc("/action/studentfeildtrip", studentFeildtrip)
  

// ~ addNew, Routes
  http.HandleFunc("/action/addNewTeacher", addNew)
  http.HandleFunc("/action/addNewStudent", addNew)
  http.HandleFunc("/action/addNewHomework", addNew)
  http.HandleFunc("/action/addNewMentor", addNew)
  http.HandleFunc("/action/addNewStaff", addNew)
  http.HandleFunc("/action/addNewProject", addNew)
  http.HandleFunc("/action/addNewMovie", addNew)
  http.HandleFunc("/action/addNewFeildtrip", addNew)
  http.HandleFunc("/action/addNewWorkshop", addNew)
  http.HandleFunc("/action/addNewSpeaker", addNew)
  http.HandleFunc("/action/addNewLecture", addNew)
  http.HandleFunc("/action/addNewQuestion", addNew)
  
  // ~ getDataList, Routes
  http.HandleFunc("/action/getDataList_Teacher", getDataList)
  http.HandleFunc("/action/getDataList_Student", getDataList)
  http.HandleFunc("/action/getDataList_Homework", getDataList)
  http.HandleFunc("/action/getDataList_Mentor", getDataList)
  http.HandleFunc("/action/getDataList_Staff", getDataList)
  http.HandleFunc("/action/getDataList_Project", getDataList)
  http.HandleFunc("/action/getDataList_Movie", getDataList)
  http.HandleFunc("/action/getDataList_Feildtrip", getDataList)
  http.HandleFunc("/action/getDataList_Workshop", getDataList)
  http.HandleFunc("/action/getDataList_Speaker", getDataList)
  http.HandleFunc("/action/getDataList_Lecture", getDataList)
  http.HandleFunc("/action/getDataList_Question", getDataList)


// . ° ~ +
 http.HandleFunc("/action/paymentInformation", paymentInformation)

// . ° ~ +
 http.HandleFunc("/homepage", indexHandler)


    // -- -
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("Loading _webapp with default port")
        }
        log.Printf("_webapp is active and Listening on port %s", port)

        log.Printf("// -- - %s", appName)
        log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal("Error Starting the HTTP Server :", err)
                return
        }




}