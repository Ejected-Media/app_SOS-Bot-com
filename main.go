 // # ~ . ~
// ,  ° . +
package main

import (
    "os"
    "log"
    
    "fmt"
		
	"text/template"
	"net/http"
	
//	"time"

)

// ,  ° . +
type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}

type pageNav struct {
    pageTitle string
    pageLink string
}

type pageList struct {
    pageTitle string
    pageLink string
}


type Todo struct {
    Title string
    Done  bool
}



type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}




// ,  ° . +
func app_welcome_center_page() {


}


// . indexHandler,  ~ for Public Pages °
func indexHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ . - // - Website App"
  pagePath := r.URL.Path
  // pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
    //  pageList: []pageNav {
  //        { pageTitle: "one", pageLink: "one"},
//          { pageTitle: "two", pageLink: "two"},
   //       { pageTitle: "three", pageLink: "three"},
    //  },
  	
  }  //. .  pageData
  
  pageHTML := "layout_main_page.html";
  if pagePath == "/" {
      pageTitle = "Index Page"
 //     pageList = pageList
  }
  
    if pagePath == "/page/classSchedule" {
      pageTitle = "classSchedule Page"
 //     pageList = pageList
 pageHTML = "/Go-Software/html-web-app/html-files/layout_sosbot_classSchedule.html";
  }


// ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler

// . testHandler,  ~ for Public Pages °
func testHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ . - // - Website App"
  pagePath := r.URL.Path
  // pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
    //  pageList: []pageNav {
  //        { pageTitle: "one", pageLink: "one"},
//          { pageTitle: "two", pageLink: "two"},
   //       { pageTitle: "three", pageLink: "three"},
    //  },
  	
  }  //. .  pageData
  
  pageHTML := "layout_main_page.html";
  if pagePath == "/" {
      pageTitle = "Index Page"
 //     pageList = pageList
  }
 


// ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, pageData)
  
}  //  .  testHandler

func hello(w http.ResponseWriter, r *http.Request) {
//	pagePath := r.URL.Path
    fmt.Fprintf(w, "Hello!")
}


func world(w http.ResponseWriter, r *http.Request) {
	pagePath := r.URL.Path
   // fmt.Fprintf(w, "World!")
   // fmt.Fprintf(w, pagePath)
    
    pageData := pagePath
    
    data := TodoPageData{
            PageTitle: pageData,
            Todos: []Todo{
                {Title: "classSchedule", Done: false},
                {Title: "studentProfiles", Done: true},
                {Title: "paymentInformation", Done: true},
                {Title: "studentAttenance", Done: true},
                {Title: "classEvent", Done: true},
                {Title: "studentHomework", Done: true},
                {Title: "studentFeildtrip", Done: true},
                {Title: "classQuestions", Done: true},
                {Title: "classMovies", Done: true},
                {Title: "studentWorkshops", Done: true},
                {Title: "listClass", Done: true},
                {Title: "loadClass", Done: true},
                {Title: "joinClass", Done: true},
                {Title: "leaveClass", Done: true},
                {Title: "addHomework", Done: true},
                {Title: "startClass", Done: true},
                {Title: "meetTeacher", Done: true},
                {Title: "askQuestions", Done: true},
                {Title: "meetStudents", Done: true},
        
            },
        }
 
 pageHTML := "layout_main_page.html";
 
  // ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, data)
  
  
}

//  .  html url routes 
//  .  as well as terminal cli logs

func main() {
// ,  ° . +
  appName := "~ . - // - Website App"
  
  http.HandleFunc("/hello", hello)
    http.HandleFunc("/world", world)
    http.HandleFunc("/page/classSchedule", world)
    http.HandleFunc("/page/studentFeildtrip", world)
    http.HandleFunc("/page/classMovies", world)

// ,  ° . +
  http.HandleFunc("/", indexHandler)
  
  // ,  ° . +
  http.HandleFunc("/one", testHandler)

// . ° ~ +
 // http.HandleFunc("/page/classSchedule", indexHandler)

// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
// ,  ° . +
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }}