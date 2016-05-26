package models

func One(c string, req interface{}, re interface{}, sel interface{}) error {
	mgoc := DB.C(c)
	return mgoc.Find(req).One(re)
}
func Update(c string, req interface{}, data interface{}) error {
	mgoc := DB.C(c)
	return mgoc.Update(req, data)
}
func Insert(c string, docs ...interface{}) error {
	mgoc := DB.C(c)
	return mgoc.Insert(docs...)
}
func List(c string, req interface{}, re interface{}, sel interface{}) {
	mgoc := DB.C(c)
	mgoc.Find(req).All(re)
}
