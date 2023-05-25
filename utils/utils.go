package utils


import (
	"strings"
	"fmt"
    "log"
	"io"
	"os"
	"encoding/base64"
	json "encoding/json" 
	xml "encoding/xml" 
	//"github.com/fsnotify/fsnotify"	// go get github.com/fsnotify/fsnotify
)

// strings.Split(url, ".")[len(strings.Split(url, "."))-1]

func CatchError() {

}
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func FileInfo(file string) map[string]interface{}{
	fileStat, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	/*
	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir() 
	*/
	return map[string]interface{}{
		"Name": fileStat.Name(),
		"Size": fileStat.Size(),
		"ModTime": fileStat.ModTime(),
		//"ModTime": fmt.Sprintf("%v", fileStat.ModTime()),
	}
}

func Fix[T any](val T) string {
	res := fmt.Sprintf("%v", val)
	res = strings.ReplaceAll(res, `/`, `\`)
	res = strings.ReplaceAll(res, `\\`, `\`)
	return res
}
func FixPath[T any](val T) string {
	res := fmt.Sprintf("%v", val)
	res = strings.ReplaceAll(res, `/`, `\`)
	res = strings.ReplaceAll(res, `\\`, `\`)
	return res
}

func Stringify[T any](devices T) string{
	res, errr := json.Marshal(devices)
	if errr == nil {
		//str := string(res)
		return string(res)
	}else {
		return "{}"
	}
}
func StringifyIndent[T any](devices T) string{
	res, errr := json.MarshalIndent(devices, "", "\t")
	if errr == nil {
		//str := string(res)
		return string(res)
	}else {
		return "{}"
	}
}

func XmlUnmarshal[T any](v T, data string) {
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		//return
	}
}
func JsonUnmarshal[T any](data string, v T) {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		//return
	}
}


func Identity[T any](v T) T {
    return v
}


/* func CatchKeyExists[T any](obj T, key string){
	_, ok := obj[key];
	return ok
} */
func ToStr[T any](v T) string{
	return fmt.Sprintf("%v", v)
}
func ConvertToString[T any](v T) string{
	return fmt.Sprintf("%v", v)
}

	func CreateDirectory(dir string) error {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("os.MkdirAll error: %v", err)
		}
		return nil
	}
	
	func CreateFile(dst string) error{
		dst = FixPath(dst)
		destination, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer destination.Close()
		return nil
	}
	func AppendToFile(filepath, data string) error{
		filepath = FixPath(filepath)
		//CreateDirectory(dir string)
		//f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			CreateFile(filepath)
			f, _ = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		}
		defer f.Close()
		_, err2 := f.WriteString("\n"+data)
		if err2 != nil {
			return err2
		}
		return nil
	}
	func WriteFile(filepath, data string) error{
		filepath = FixPath(filepath)
		f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			//return ""
			//log.Fatalf("unable to read file: %v", err)
			err = CreateFile(filepath)
			if(err != nil){
				fmt.Println("utils.WriteFile.CreateFile ERROR")
				return err
			}
			f, _ = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		}
		defer f.Close()

		_, err2 := f.WriteString(data)
		if err2 != nil {
			fmt.Println("utils.WriteFile.WriteString ERROR")
			return err2
			//log.Fatal(err2)
		}
		fmt.Println("done")
		return nil
	}

	func Copy(src, dst string) (int64, error) {
		sourceFileStat, err := os.Stat(src)
		if err != nil {
				return 0, err
		}
	
		if !sourceFileStat.Mode().IsRegular() {
				return 0, fmt.Errorf("%s is not a regular file", src)
		}
	
		source, err := os.Open(src)
		if err != nil {
				return 0, err
		}
		defer source.Close()
	
		destination, err := os.Create(dst)
		if err != nil {
				return 0, err
		}
		defer destination.Close()
		nBytes, err := io.Copy(destination, source)
		return nBytes, err
	}

	func Delete(src string) error {
		e := os.Remove(src)
		if e != nil {
			log.Fatal(e)
		}
		return nil
	}
	
	func File(filename string) string{
		filename = FixPath(filename)
		//f, err := os.Open("file.txt")
		f, err := os.Open(filename)
		if err != nil {
			return ""
			log.Fatalf("unable to read file: %v", err)
		}
		defer f.Close()
		buf := make([]byte, 1024)
		str := ""
		for {
			n, err := f.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("FILE ", err)
				log.Fatalf("unable to read file: %v", err)
				continue
			}
			if n > 0 {
				str += string(buf[:n])
			}
		}
		return str
	}

	func PathExists(path string) (bool, error) {
		_, err := os.Stat(path)
		if err == nil { return true, nil }
		if os.IsNotExist(err) { return false, nil }
		return false, err
	}