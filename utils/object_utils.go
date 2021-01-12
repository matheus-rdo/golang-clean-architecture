package utils

//CopyObject retorna uma nova cópia de um objeto
func CopyObject(obj map[string]interface{}) map[string]interface{} {
	if obj == nil {
		return nil
	}
	copiedObj := make(map[string]interface{})

	/* Copy Content from originObj to copiedObj*/
	for index, element := range obj {
		copiedObj[index] = element
	}

	return copiedObj
}
