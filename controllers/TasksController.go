package TaskController

import (
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"project_golang/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	
)

func sqliteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"),&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		
	}

	return db
}

func renderTemplateHTML(htmlTmp string,w http.ResponseWriter, data interface{}){
	files := [] string{
		"views/"+htmlTmp+".html",
		"views/base.html",
		
	}

	tmpt, err := template.ParseFiles(files...)
	if err != nil {
		panic("Error parsing : "+ err.Error())

	}
	errExec := tmpt.ExecuteTemplate(w,"base",data)

	if errExec != nil {
		panic("Error Execute: "+ errExec.Error())
	}
}

func Index(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	db := sqliteDB()

	var task_model []Model.Task_model
	db.Find(&task_model)
	list_data := map[string]interface{}{
		"Task_model":task_model,
	}
	
	
	renderTemplateHTML("index", w, list_data)
}

func Create(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	db := sqliteDB()
	if r.Method == "POST"{
		task_data := Model.Task_model{
			Nama : r.FormValue("nama"),
			Pegawai: r.FormValue("pegawai"),
			Status: r.FormValue("status"),
			Tgl : r.FormValue("tgl"),
		}
		db.Create(&task_data)
		http.Redirect(w,r,"/",http.StatusFound)
		//  taskdata := r.FormValue("nama")
		// fmt.Println(taskdata)
		// renderTemplateHTML("index", w, nil)
	}else{

		renderTemplateHTML("create", w, nil)
	}
}

func Update(w http.ResponseWriter, r * http.Request,params httprouter.Params){
	db:= sqliteDB()
	list_data_model := Model.Task_model{}
	err := db.First(&list_data_model,params.ByName("id")).Error
	if err != nil {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	if r.Method == "POST"{
			list_data_model.Nama = r.FormValue("nama")
			list_data_model.Pegawai= r.FormValue("pegawai")
			list_data_model.Status= r.FormValue("status")
			list_data_model.Tgl = r.FormValue("tgl")
		
		db.Save(&list_data_model)
		http.Redirect(w,r,"/",http.StatusFound)
	
	}else{
		list_data_model := Model.Task_model{}
		db.First(&list_data_model,params.ByName("id"))
		list_data := map[string]interface{}{
			"Task_model":list_data_model,
		}
		renderTemplateHTML("update",w,list_data)
		
	}

	

}

func DeleteTask(w http.ResponseWriter, r * http.Request,params httprouter.Params){
	db:= sqliteDB()
	list_data_model := Model.Task_model{}
	err := db.First(&list_data_model,params.ByName("id")).Error
	if err != nil {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	db.Delete(&list_data_model,params.ByName("id"))
	http.Redirect(w,r,"/",http.StatusFound)
}
